/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FromFile(path string) (*Genome, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	g, err := ReadGenome(file)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func ReadGenome(r io.ReadCloser) (*Genome, error) {
	g := &Genome{
		Nodes: make([]*Node, 0),
		Genes: make([]*Gene, 0),
	}
	var err error

	defer r.Close()

	sc := bufio.NewScanner(r)

	discardCommentRegexp := regexp.MustCompile(`^(\s*/{2,}|/{2,})`)

	for sc.Scan() {
		line := sc.Text()

		err = parseLine(discardCommentRegexp, line, g)
		if err != nil {
			return nil, err
		}
	}

	if err = sc.Err(); err == nil && len(g.Genes) == 0 && len(g.Nodes) == 0 {
		return nil, ErrEmptyGenomeFile
	}

	return g, nil
}

func parseLine(discardCommentRegexp *regexp.Regexp, line string, g *Genome) error {
	if len(discardCommentRegexp.FindStringSubmatch(line)) == 0 {

		if err := parseGenomeId(g, line); err != nil {
			return err
		}

		if err := parseNode(g, line); err != nil {
			return err
		}

		if err := parseGene(g, line); err != nil {
			return err
		}
	}
	return nil
}

func parseGenomeId(g *Genome, line string) error {

	if strings.Contains(line, "genomestart") {
		str := strings.Split(line, " ")

		if id, err := strconv.ParseInt(str[1], 10, 64); err != nil {
			return err
		} else {
			g.Id = uint(id)
		}
	}

	return nil
}

func parseNode(g *Genome, line string) error {

	if strings.Contains(line, "node") {

		str := strings.Split(line, " ")

		n := &Node{}

		id, err := strconv.ParseUint(str[1], 10, 64)
		if err != nil {
			return err
		}

		n.Id = uint(id)

		nodeType, err := strconv.ParseUint(str[2], 10, 64)
		if err != nil {
			return err
		}

		n.NodeType = NodeType(nodeType)

		g.Nodes = append(g.Nodes, n)
	}

	return nil
}

func parseGene(g *Genome, line string) error {
	if strings.Contains(line, "gene") {
		// innovation/id, weight, input id, output id, enabled/disabled
		str := strings.Split(line, " ")
		innovation, err := strconv.ParseUint(str[1], 10, 64)
		if err != nil {
			return err
		}

		weight, err := strconv.ParseFloat(str[2], 64)
		if err != nil {
			return err
		}

		inId, err := strconv.ParseUint(str[3], 10, 64)
		if err != nil {
			return err
		}

		outId, err := strconv.ParseUint(str[4], 10, 64)
		if err != nil {
			return err
		}

		enabledNum, err := strconv.ParseInt(str[5], 10, 64)
		if err != nil {
			return err
		}

		intToBool := func(v int) bool {
			return v >= 1
		}

		isEnabled := intToBool(int(enabledNum))

		// TODO: Implement method to get a Node by id from Genome
		g.Genes = append(g.Genes, &Gene{Innovation: uint(innovation), Weight: weight, InNode: g.Nodes.fetch(uint(inId)), OutNode: g.Nodes.fetch(uint(outId)), Enabled: isEnabled, Recurrent: false})
	}

	return nil
}
