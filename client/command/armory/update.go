package armory

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"io/ioutil"
	"strings"

	"github.com/bishopfox/sliver/client/assets"
	"github.com/bishopfox/sliver/client/command/extensions"
	"github.com/bishopfox/sliver/client/console"
	"github.com/desertbit/grumble"
)

// ArmoryUpdateCmd - Update all installed extensions/aliases
func ArmoryUpdateCmd(ctx *grumble.Context, con *console.SliverConsoleClient) {
	con.PrintInfof("Refreshing package cache ... ")
	clientConfig := parseArmoryHTTPConfig(ctx)
	refresh(clientConfig)
	con.Printf("done!\n")
	updates := checkForExtensionUpdates(clientConfig, con)
	if 0 < len(updates) {
		con.PrintInfof("%d extension(s) out of date: %s\n", len(updates), strings.Join(updates, ", "))
	} else {
		con.PrintInfof("All extensions up to date!\n")
	}
}

func checkForExtensionUpdates(clientConfig ArmoryHTTPConfig, con *console.SliverConsoleClient) []string {
	_, cachedExtensions := packagesInCache()
	results := []string{}
	for _, extManifestPath := range assets.GetInstalledExtensionManifests() {
		data, err := ioutil.ReadFile(extManifestPath)
		if err != nil {
			continue
		}
		localManifest, err := extensions.ParseExtensionManifest(data)
		if err != nil {
			continue
		}
		for _, latestExt := range cachedExtensions {
			// Right now we don't try to enforce any kind of versioning, it is assumed if the version from
			// the armory differs at all from the local version, the extension is out of date.
			if latestExt.CommandName == localManifest.CommandName && latestExt.Version != localManifest.Version {
				results = append(results, localManifest.CommandName)
			}
		}
	}
	return results
}
