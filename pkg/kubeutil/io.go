// Copyright 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

package kubeutil

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ListGoFiles returns a list of all go files in the root directory
func ListGoFiles(root string) ([]string, error) {
	return listFiles(root, []string{".go"})
}

// ListYAMLFiles returns a list of all yaml files in the root directory
func ListYAMLFiles(root string) ([]string, error) {
	return listFiles(root, []string{".yaml", ".yml"})
}

func listFiles(root string, extensions []string) ([]string, error) {
	var files []string

	fi, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.New("root is not a directory")
	}
	err = filepath.Walk(
		root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("walk %q %q, %w", path, info.Name(), err)
			}

			if !info.IsDir() && contains(
				filepath.Ext(filepath.Base(path)),
				extensions,
			) {
				files = append(files, path)
			}

			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("walk: %w", err)
	}
	return files, nil
}

func contains(e string, s []string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ReadManifest reads a YAML file and splits it into a list of YAML documents
func ReadManifest(filePath string) ([]string, error) {
	e := filepath.Ext(filePath)
	if e != ".yaml" && e != ".yml" {
		return nil, fmt.Errorf("not yaml file: %s", filePath)
	}
	yf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read manifest %s: %w", filePath, err)
	}
	splitYaml, err := ManifestSplit(bytes.NewReader(yf))
	if err != nil {
		return nil, fmt.Errorf("splitting manifest: %s: %w", filePath, err)
	}
	return splitYaml, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}