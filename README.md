Unicode transliterator (also known as unidecode) for Go
=======================================================

Use the following command to install gounidecode

go get http://github.com/fiam/gounidecode/unidecode

Example usage
=============

package main

import (
    "fmt"
    "github.com/fiam/gounidecode/unidecode"
)

func main() {
    fmt.Println(Unidecode("áéíóú")) // Will print aeiou
    fmt.Println(Unidecode("\u5317\u4EB0")) // Will print Bei Jing
    fmt.Println(Unidecode("Κνωσός")) // Will print Knosos
}
