package assets

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

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
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	ver "github.com/bishopfox/sliver/client/version"
)

var (
	//go:embed fs/*
	clientAssetsFs embed.FS
)

const (
	// SliverClientDirName - Directory storing all of the client configs/logs
	SliverClientDirName = ".sliver-client"

	versionFileName = "version"
)

// GetRootAppDir - Get the Sliver app dir ~/.sliver-client/
func GetRootAppDir() string {
	user, _ := user.Current()
	dir := filepath.Join(user.HomeDir, SliverClientDirName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func assetVersion() string {
	appDir := GetRootAppDir()
	data, err := ioutil.ReadFile(filepath.Join(appDir, versionFileName))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

func saveAssetVersion(appDir string) {
	versionFilePath := filepath.Join(appDir, versionFileName)
	fVer, _ := os.Create(versionFilePath)
	defer fVer.Close()
	fVer.Write([]byte(ver.GitCommit))
}

// Setup - Extract or create local assets
func Setup(force bool, echo bool) {
	appDir := GetRootAppDir()
	localVer := assetVersion()
	if force || localVer == "" || localVer != ver.GitCommit {
		if echo {
			fmt.Printf("Unpacking assets ...\n")
		}
		err := setupDefaultExtensions(appDir)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		saveAssetVersion(appDir)
	}
	if _, err := os.Stat(filepath.Join(appDir, settingsFileName)); os.IsNotExist(err) {
		SaveSettings(nil)
	}
}

func setupDefaultExtensions(appDir string) error {
	localExtDir := GetExtensionsDir()
	rootEmbedPath := "fs/extensions"
	return fs.WalkDir(clientAssetsFs, rootEmbedPath, func(embedPath string, embedDir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		localPath := strings.TrimPrefix(embedPath, rootEmbedPath)
		if embedDir.IsDir() {
			err := os.MkdirAll(filepath.Join(localExtDir, localPath), 0o700)
			if err != nil {
				return err
			}
		} else {
			data, err := fs.ReadFile(clientAssetsFs, embedPath)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(localExtDir, localPath), data, 0o600)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
