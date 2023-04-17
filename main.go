package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	path := "/Users/robsonpetinari/Documents/EFD_00029_00001.txt"
	arquivoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	arquivo := strings.Split(string(arquivoBytes), "\n")
	_0200List := []string{}
	_C170List := []string{}
	erros := 0

	for _, line := range arquivo {
		if strings.Contains(line, "|0200|") {
			_0200List = append(_0200List, line)
		}
		if strings.Contains(line, "|C170|") {
			_C170List = append(_C170List, line)
		}
	}

	for _, item0200 := range _0200List {
		linha0200 := strings.Split(item0200, "|")
		for _, itemC170 := range _C170List {
			linhaC170 := strings.Split(itemC170, "|")
			if linha0200[2] == linhaC170[3] {
				if linha0200[6] != linhaC170[6] {
					linhaC170Novo := append([]string(nil), linhaC170...)
					linhaC170Novo[6] = linha0200[6]
					for indexC170, line := range arquivo {
						if line == itemC170 {
							arquivo[indexC170] = strings.Join(linhaC170Novo, "|")
							break
						}
					}
					erros++
				}
			}
		}
	}

	fmt.Println(erros)
	err = ioutil.WriteFile(path, []byte(strings.Join(arquivo, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}
