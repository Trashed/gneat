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
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	discardCommentdRegexp := regexp.MustCompile(`^(\s*\/{2,}|\/{2,})`)

	for sc.Scan() {
		line := sc.Text()

		if len(discardCommentdRegexp.FindStringSubmatch(line)) == 0 {

			if strings.Contains(line, "genomestart") {
				if err = parseGenomeId(g, line); err != nil {
					return nil, err
				}
			}

			if strings.Contains(line, "nodetrait") {
				if err = parseNodeTrait(g, line); err != nil {
					return nil, err
				}
			}
		}
	}

	return g, nil
}

func parseGenomeId(g *Genome, line string) error {
	str := strings.Split(line, " ")

	if id, err := strconv.ParseInt(str[1], 10, 64); err != nil {
		return err
	} else {
		g.Id = uint(id)
	}

	return nil
}

func parseNodeTrait(g *Genome, line string) error {
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

	return nil
}
