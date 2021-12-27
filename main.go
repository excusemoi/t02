package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	pathKey       = "path"
	wordKey       = "word"
	outputPathKey = "outputPath"
)

var enVarietyTranslations = [][][]string{
	[][]string{
		[]string{"Ya"},
		[]string{"Я"},
	},
	[][]string{
		[]string{"Ja"},
		[]string{"я"},
	},
	[][]string{
		[]string{"Je"},
		[]string{"Э"},
	},
	[][]string{
		[]string{"je"},
		[]string{"э"},
	},
	[][]string{
		[]string{"Ju"},
		[]string{"Ю"},
	},
	[][]string{
		[]string{"ju"},
		[]string{"ю"},
	},
	[][]string{
		[]string{"Yu"},
		[]string{"Ю"},
	},
	[][]string{
		[]string{"yu"},
		[]string{"ю"},
	},
	[][]string{
		[]string{"Ch"},
		[]string{"Ч"},
	},
	[][]string{
		[]string{"ch"},
		[]string{"ч"},
	},
	[][]string{
		[]string{"Ch"},
		[]string{"Ч"},
	},
	[][]string{
		[]string{"Shh"},
		[]string{"Щ"},
	},
	[][]string{
		[]string{"shh"},
		[]string{"щ"},
	},
	[][]string{
		[]string{"Sh"},
		[]string{"ш"},
	},
	[][]string{
		[]string{"sh"},
		[]string{"ш"},
	},
	[][]string{
		[]string{"Zh"},
		[]string{"Ж"},
	},
	[][]string{
		[]string{"zh"},
		[]string{"ж"},
	},
	[][]string{
		[]string{"Jo"},
		[]string{"Ё"},
	},
	[][]string{
		[]string{"jo"},
		[]string{"ё"},
	},
	[][]string{
		[]string{"Yo"},
		[]string{"Ё"},
	},
	[][]string{
		[]string{"yo"},
		[]string{"ё"},
	},
	[][]string{
		[]string{"H"},
		[]string{"Х"},
	},
	[][]string{
		[]string{"h"},
		[]string{"х"},
	},
	[][]string{
		[]string{"X"},
		[]string{"Кс"},
	},
	[][]string{
		[]string{"x"},
		[]string{"кс"},
	},
	[][]string{
		[]string{"A"},
		[]string{"А"},
	},
	[][]string{
		[]string{"B"},
		[]string{"Б"},
	},
	[][]string{
		[]string{"a"},
		[]string{"а"},
	},
	[][]string{
		[]string{"b"},
		[]string{"б"},
	},
	[][]string{
		[]string{"C"},
		[]string{"С"},
	},
	[][]string{
		[]string{"c"},
		[]string{"с"},
	},
	[][]string{
		[]string{"D"},
		[]string{"Д"},
	},
	[][]string{
		[]string{"d"},
		[]string{"д"},
	},
	[][]string{
		[]string{"E"},
		[]string{"Е"},
	},
	[][]string{
		[]string{"e"},
		[]string{"е"},
	},
	[][]string{
		[]string{"f"},
		[]string{"ф"},
	},
	[][]string{
		[]string{"F"},
		[]string{"Ф"},
	},
	[][]string{
		[]string{"I"},
		[]string{"И"},
	},
	[][]string{
		[]string{"i"},
		[]string{"и"},
	},
	[][]string{
		[]string{"G"},
		[]string{"Г"},
	},
	[][]string{
		[]string{"g"},
		[]string{"г"},
	},
	[][]string{
		[]string{"k"},
		[]string{"к"},
	},
	[][]string{
		[]string{"K"},
		[]string{"К"},
	},
	[][]string{
		[]string{"L"},
		[]string{"Л"},
	},
	[][]string{
		[]string{"J"},
		[]string{"Й"},
	},
	[][]string{
		[]string{"j"},
		[]string{"й"},
	},
	[][]string{
		[]string{"l"},
		[]string{"л"},
	},
	[][]string{
		[]string{"M"},
		[]string{"М"},
	},
	[][]string{
		[]string{"m"},
		[]string{"м"},
	},
	[][]string{
		[]string{"n"},
		[]string{"н"},
	},
	[][]string{
		[]string{"O"},
		[]string{"О"},
	},
	[][]string{
		[]string{"o"},
		[]string{"о"},
	},
	[][]string{
		[]string{"P"},
		[]string{"П"},
	},
	[][]string{
		[]string{"p"},
		[]string{"п"},
	},
	[][]string{
		[]string{"Q"},
		[]string{"К"},
	},
	[][]string{
		[]string{"q"},
		[]string{"к"},
	},
	[][]string{
		[]string{"R"},
		[]string{"Р"},
	},
	[][]string{
		[]string{"S"},
		[]string{"С"},
	},
	[][]string{
		[]string{"s"},
		[]string{"с"},
	},
	[][]string{
		[]string{"T"},
		[]string{"Т"},
	},
	[][]string{
		[]string{"t"},
		[]string{"т"},
	},
	[][]string{
		[]string{"U"},
		[]string{"У"},
	},
	[][]string{
		[]string{"u"},
		[]string{"у"},
	},
	[][]string{
		[]string{"V"},
		[]string{"В"},
	},
	[][]string{
		[]string{"v"},
		[]string{"в"},
	},
	[][]string{
		[]string{"W"},
		[]string{"В"},
	},
	[][]string{
		[]string{"w"},
		[]string{"в"},
	},
	[][]string{
		[]string{"X"},
		[]string{"Кс"},
	},
	[][]string{
		[]string{"x"},
		[]string{"кс"},
	},
	[][]string{
		[]string{"Y"},
		[]string{"Ы"},
	},
	[][]string{
		[]string{"y"},
		[]string{"ы"},
	},
	[][]string{
		[]string{"Z"},
		[]string{"З"},
	},
	[][]string{
		[]string{"z"},
		[]string{"з"},
	},
	[][]string{
		[]string{"N"},
		[]string{"Н"},
	},
}

var enTranslations = map[string]string{
	"Ya|Ja": "Я",
	"ya|ja": "я",
	"Je":    "Э",
	"je":    "э",
	"Ju|Yu": "Ю",
	"ju|yu": "ю",
	"Ch":    "Ч",
	"ch":    "ч",
	"Shh|W": "Щ",
	"shh|w": "щ",
	"Sh":    "Ш",
	"sh":    "ш",
	"Zh":    "Ж",
	"zh":    "ж",
	"Yo|Jo": "Ё",
	"yo|jo": "ё",
	"H":     "Х",
	"h":     "х",
	"X":     "Кс",
	"x":     "кс",
	"##":    "ъ",
	"A":     "А",
	"a":     "а",
	"B":     "Б",
	"b":     "б",
	"V":     "В",
	"v":     "в",
	"G":     "Г",
	"g":     "г",
	"D":     "Д",
	"d":     "д",
	"Z":     "З",
	"z":     "з",
	"I":     "И",
	"i":     "и",
	"J":     "Й",
	"j":     "й",
	"K":     "К",
	"k":     "к",
	"L":     "Л",
	"l":     "л",
	"M":     "М",
	"m":     "м",
	"N":     "Н",
	"n":     "н",
	"O":     "О",
	"o":     "о",
	"P":     "П",
	"p":     "п",
	"R":     "Р",
	"r":     "р",
	"S":     "С",
	"s":     "с",
	"T":     "Т",
	"t":     "т",
	"U":     "У",
	"u":     "у",
	"F":     "Ф",
	"f":     "ф",
	"C":     "Ц",
	"c":     "ц",
	"Y":     "Ы",
	"y":     "ы",
	"#":     "ъ",
	"E":     "Е",
	"e":     "е",
}

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
	for k, v := range enTranslations {
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
