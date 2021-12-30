package main

import (
	"flag"
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

func VariableTranslit(s string) map[string]struct{} {
	var (
		cyrillicStrings        = make(map[string]struct{})
		cyrillicString  string = s
	)
	for i, arr := range translations.EnVarietyTranslations {
		r, _ := regexp.Compile(arr[0][0])
		cyrillicString = r.ReplaceAllString(cyrillicString, arr[1][0])
		if len(arr[1]) > 1 {
			for _, ch := range arr[1] {
				cyrillicStringCp := s
				for j, nArr := range translations.EnVarietyTranslations {
					r, _ = regexp.Compile(nArr[0][0])
					if j == i {
						cyrillicStringCp = r.ReplaceAllString(cyrillicStringCp, ch)
					} else {
						cyrillicStringCp = r.ReplaceAllString(cyrillicStringCp, nArr[1][0])
					}
				}
				cyrillicStrings[cyrillicStringCp] = struct{}{}
			}
		}
	}
	cyrillicStrings[cyrillicString] = struct{}{}
	return cyrillicStrings
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
	}
	translited := VariableTranslit(contentToTranslit)
	contentToTranslit = ""
	for k := range translited {
		contentToTranslit += k + "\n"
	}
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
