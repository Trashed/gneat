/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genome

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Trashed/gneat/activation"
)

func FromFile(path string) (*Genome, error) {
	g := &Genome{
		NodeTraits: make([]*NodeTrait, 0),
		Nodes:      make([]*Node, 0),
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	discardCommentRegexp := regexp.MustCompile(`^(\s*/{2,}|/{2,})`)

	for sc.Scan() {
		line := sc.Text()

		genome, err2 := parseLine(discardCommentRegexp, line, err, g)
		if err2 != nil {
			return genome, err2
		}
	}

	return g, nil
}

func parseLine(discardCommentRegexp *regexp.Regexp, line string, err error, g *Genome) (*Genome, error) {
	if len(discardCommentRegexp.FindStringSubmatch(line)) == 0 {

		if err = parseGenomeId(g, line); err != nil {
			return nil, err
		}

		if err = parseNodeTrait(g, line); err != nil {
			return nil, err
		}

		if err = parseNode(g, line); err != nil {
			return nil, err
		}
	}
	return nil, nil
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

func parseNodeTrait(g *Genome, line string) error {

	if strings.Contains(line, "nodetrait") {

		str := strings.Split(line, " ")

		t := &NodeTrait{}

		traitId, err := strconv.ParseUint(str[1], 10, 64)
		if err != nil {
			return err
		}

		t.Id = uint(traitId)

		actFn, err := strconv.ParseInt(str[2], 10, 64)
		if err != nil {
			return err
		}

		t.ActivationFunc = activation.ActivationFuncEnum(actFn)

		bias, err := strconv.ParseFloat(str[3], 64)
		if err != nil {
			return err
		}

		t.Bias = bias

		g.NodeTraits = append(g.NodeTraits, t)
	}

	return nil
}

func parseNode(g *Genome, line string) error {

	if strings.Contains(line, "node") && !strings.Contains(line, "nodetrait") {

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

		traitId, err := strconv.ParseUint(str[3], 10, 64)
		if err != nil {
			return err
		}

		n.Trait = getNodeTrait(g.NodeTraits, uint(traitId))

		g.Nodes = append(g.Nodes, n)
	}

	return nil
}
