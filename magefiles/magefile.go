// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

//go:build mage
// +build mage

package main

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/sh"
	"github.com/princjef/mageutil/bintool"
	"github.com/princjef/mageutil/shellcmd"
)

var Default = All

var (
	linter = bintool.Must(bintool.New(
		"golangci-lint{{.BinExt}}",
		"1.51.1",
		"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
	))

	archmap = map[string]string{
		"386":   "x86_32",
		"amd64": "x86_64",
		"arm64": "aarch_64",
	}

	protoc *bintool.BinTool

	protocGenGo = bintool.Must(bintool.NewGo(
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		"v1.28.1",
	))

	protocGenGoGrpc = bintool.Must(bintool.NewGo(
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc",
		"v1.3.0",
	))
)

// crazy....
func init() {
	suffix := runtime.GOOS
	if suffix == "windows" {
		if runtime.GOARCH == "386" {
			suffix = "win32"
		} else {
			suffix = "win64"
		}
	} else {
		if suffix == "darwin" {
			suffix = "osx"
		}

		arch := "x86_64"
		if runtime.GOARCH == "386" {
			arch = "x86_32"
		} else if runtime.GOARCH == "arm64" {
			arch = "aarch_64"
		}

		suffix = suffix + "-" + arch
	}

	protoc = bintool.Must(bintool.New(
		"protoc{{.BinExt}}",
		"23.3",
		"https://github.com/protocolbuffers/protobuf/releases/download/v{{.Version}}/protoc-{{.Version}}-"+suffix+"{{.ArchiveExt}}",
		bintool.WithArchiveExt(".zip"),
	))
}

func Lint() error {
	if err := linter.Ensure(); err != nil {
		return err
	}

	return linter.Command(`run`).Run()
}

func Generate() error {
	if err := protoc.Ensure(); err != nil {
		return err
	}

	if err := protocGenGo.Ensure(); err != nil {
		return err
	}

	if err := protocGenGoGrpc.Ensure(); err != nil {
		return err
	}

	err := protoc.Command(`--plugin=protoc-gen-go=bin/protoc-gen-go --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc -I . --go_opt=module=github.com/szkiba/xk6-output-plugin-go --go-grpc_opt=module=github.com/szkiba/xk6-output-plugin-go --go-grpc_out=. --go_out=. proto/output.proto`).Run()
	if err != nil {
		return err
	}

	return License()
}

func glob(patterns ...string) (string, error) {
	buff := new(strings.Builder)

	for _, p := range patterns {
		m, err := filepath.Glob(p)
		if err != nil {
			return "", err
		}

		_, err = buff.WriteString(strings.Join(m, " ") + " ")
		if err != nil {
			return "", err
		}
	}

	return buff.String(), nil
}

func License() error {
	all, err := glob(
		"output/*.go",
		".*.yml",
		".gitignore",
		"magefiles/*go",
		"magefiles/.gitignore",
	)
	if err != nil {
		return err
	}

	return shellcmd.Command(
		`reuse annotate --copyright "Iván Szkiba" --merge-copyrights --license MIT --skip-unrecognised ` + all,
	).Run()
}

func Clean() error {
	sh.Rm("magefiles/bin")
	sh.Rm("coverage.txt")
	sh.Rm("bin")

	return nil
}

func All() error {
	return Lint()
}
