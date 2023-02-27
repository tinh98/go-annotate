package main

import (
	"flag"
	"fmt"
	"go-annotate/cmd"
	"go-annotate/pkg/goan"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	pkgName   string
	fileName  string
	toolType  string
	outputDir string
)

func init() {
	flag.StringVar(&pkgName, "p", "", "The package name")
	flag.StringVar(&fileName, "f", "", "The file name")
	flag.StringVar(&toolType, "t", "swagger", "The tool type (default: swagger)")
	flag.StringVar(&outputDir, "o", "", "The output directory")
}

func main() {
	cmd.Execute()
	flag.Parse()

	if pkgName == "" || fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	args := []string{
		"go-annotate",
		fmt.Sprintf("-p %s", pkgName),
		fmt.Sprintf("-f %s", fileName),
		fmt.Sprintf("-t %s", toolType),
	}

	if outputDir != "" {
		args = append(args, fmt.Sprintf("-o %s", outputDir))
	}

	cmd := exec.Command("go", "generate", strings.Join(args, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running go-annotate: %s\n", err)
		os.Exit(1)
	}

	goan.GenerateSwaggerAnnotations(path.Join(fileName, pkgName))
}
