/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genome

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FromFile(path string) (*Genome, error) {
	g := &Genome{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, "genomestart") {
			if err = initGenome(g, line); err != nil {
				return nil, err
			}
		}
	}

	return g, nil
}

func initGenome(g *Genome, line string) error {
	str := strings.Split(line, " ")

	if id, err := strconv.ParseInt(str[1], 10, 64); err != nil {
		return err
	} else {
		g.Id = int(id)
	}

	return nil
}
