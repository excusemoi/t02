package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"search/t02/translations"
)

const (
	pathKey       = "path"
	wordKey       = "word"
	outputPathKey = "outputPath"
)

func parseCommandLineArgs() map[string]string {
	var (
		path       = flag.String("p", "", "path to file with words to translit")
		word       = flag.String("w", "", "word to translit")
		outputPath = flag.String("o", "", "path to translitered words")
		args       = make(map[string]string)
	)
	flag.Parse()
	if *path != "" {
		args[pathKey] = *path
	}
	if *word != "" {
		args[wordKey] = *word
	}
	if *outputPath != "" {
		args[outputPathKey] = *outputPath
	}
	return args
}

func VariableTranslit(s string) string {
	return ""
}

func Translit(s string) string {
	cyrillicString := s
	for k, v := range translations.EnTranslations {
		r, _ := regexp.Compile(k)
		cyrillicString = r.ReplaceAllString(cyrillicString, v)
	}
	return cyrillicString
}

func TranslitUtil() {
	var (
		args              = parseCommandLineArgs()
		contentToTranslit string
	)
	if word, in := args[wordKey]; in {
		contentToTranslit += word + "\n"
	}
	if path, in := args[pathKey]; in {
		fileContent, err := getFileContent(path)
		if err != nil {
			io.WriteString(os.Stderr, err.Error()+" ")
			os.Exit(1)
		}
		contentToTranslit += fileContent
		fmt.Println(contentToTranslit)
	}
	contentToTranslit = Translit(contentToTranslit)
	if outputPath, in := args[outputPathKey]; in {
		if err := writeContentToFile(outputPath, contentToTranslit); err != nil {
			io.WriteString(os.Stderr, err.Error()+" ")
			os.Exit(1)
		}
	}
}

func writeContentToFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func getFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	TranslitUtil()
}
