/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Trashed/gneat"
)

func FromFile(path string) gneat.NeatConfig {

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error while reading the configuration file: %v\n", err)
		return gneat.NeatConfig{}
	}

	conf := gneat.NeatConfig{}

	err = json.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("error while parsing file content to NeatConfig struct: %v\n", err)
		return conf
	}

	return conf
}
