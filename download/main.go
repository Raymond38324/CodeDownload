package main

import "fmt"

func main() {
	if err := DownloadFile("http://idea.medeming.com/jets/images/jihuoma.zip", ".code.zip"); err != nil{
		fmt.Println(err)
		return
	}
	if code, err := ExtraFile(".code.zip"); err != nil {
		panic(err)
	} else {
		Exec(code)
		fmt.Println("按 Command + v 粘贴")
	}
}