package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"search/t02/translations"
	"strings"
	"unicode"

	"github.com/gen1us2k/transcript"
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
	s = strings.TrimSpace(s)
	var (
		cyrillicStrings        = make(map[string]struct{})
		cyrillicString  string = s
	)
	for i, arr := range translations.EnVarietyTranslations {
		r, _ := regexp.Compile(arr[0][0])
		cyrillicString = ReplaceAllString(r, cyrillicString, arr[1][0])
		if len(arr[1]) > 1 {
			for _, ch := range arr[1] {
				cyrillicStringCp := s
				for j, nArr := range translations.EnVarietyTranslations {
					r, _ = regexp.Compile(nArr[0][0])
					if j == i {
						cyrillicStringCp = ReplaceAllString(r, cyrillicStringCp, ch)
					} else {
						cyrillicStringCp = ReplaceAllString(r, cyrillicStringCp, nArr[1][0])
					}
				}
				cyrillicStrings[cyrillicStringCp] = struct{}{}
			}
		}
	}
	cyrillicStrings[cyrillicString] = struct{}{}
	return cyrillicStrings
}

func TransciptionTranslit(s string) map[string]struct{} {
	cyrillicStrings := make(map[string]struct{})
	cyrillicStrings[transcript.TransliterateRussian(s)] = struct{}{}
	return cyrillicStrings
}

func ReplaceAllString(r *regexp.Regexp, src, repl string) string {
	//sr := []rune(strings.TrimSpace(src))
	/*if sr[0] == 'h' {
		src = src[1:]
		src = "эйч" + src
	}
	if sr[0] == 'H' {
		src = src[1:]
		src = "Эйч" + src
	}
	if sr[0] == 'j' {
		src = src[1:]
		src = "джи" + src
	}
	*/
	/*if sr[len(sr)-1] == 'p' {
		src = string(sr[:len(sr)-1])
		src += "пи"
		return src
	}*/
	/*
		if sr[len(sr)-1] == 'P' {
			src = string(sr[:len(sr)-1])
			src += "Пи"
		}*/
	return r.ReplaceAllString(src, repl)
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
	fmt.Println(contentToTranslit)
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

func TranslitCsv(pathToFile string) error {
	var (
		file, err  = os.Open(pathToFile)
		newFile, _ = os.Create("translited" + pathToFile)
	)
	if err != nil {
		return err
	}
	defer file.Close()
	defer newFile.Close()

	reader := csv.NewReader(file)
	writer := csv.NewWriter(newFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = 4
	reader.LazyQuotes = true

	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		if containsCyrillic(record[2]) {
			continue
		}
		translited := TransciptionTranslit(strings.ToLower(record[2]))
		translitedString := ""
		for k := range translited {
			translitedString += k + "|"
		}
		fmt.Printf("%s translited -> %s\n", record[2], translitedString)
		writer.Write([]string{record[2], translitedString})
	}
	return nil
}

func containsCyrillic(s string) bool {
	sr := []rune(s)
	for _, r := range sr {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}

func main() {
	transcript.LoadDict("")
	fmt.Println(transcript.GetTranscription("абв"))
}
