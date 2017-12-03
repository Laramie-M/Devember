package main


import (
    "fmt"
    //"io"
    "os"
)

func errorCheck(e error){
	if e != nil{
		panic(e)
	}
}

func fileOpen(relPath string) *os.File {
	pwd, _ := os.Getwd()
	f, e1 := os.Open(pwd + relPath)
	errorCheck(e1)
	return f
}

func typeCheck(head []byte) {
	//file Types
	pdf := []byte{37, 80, 68, 70}
	gif := []byte{71, 73, 70, 76}
	jpg := []byte{255,216,255}
	bmp := []byte{66,77}
	png := []byte{137, 80, 78, 71}
	ppm := []byte{80,54}
	riff := []byte{82,73,70,70}
	bpg := []byte{66, 80, 71, 251}
	//master switch checks head[0]
	switch head[0]{
	case 66:
		if head[1] == bmp[1]{fmt.Println("BMP")}else
		if head[2] == bpg[2] {fmt.Println("BPG")}
	case 71:
		if head[1] == gif[1] && head[3] == gif[3]{fmt.Println("GIF")}
	case 255:
		if head[1] == jpg[1] && head[2] == jpg[2]{fmt.Println("JPG")}
	case 37:
		if head[1] ==  pdf[1] && head[2] == pdf[2]{fmt.Println("PDF")}
	case 80:
		if head[1] ==  ppm[1] {fmt.Println("PPM")}
	case 137:
		if head[1] ==  png[1] && head[2] == png[2]{fmt.Println("PNG")}
	case 82:
		if head[1] ==  riff[1] && head[2] == riff[2]{fmt.Println("WAV or AVI or WebP")}

	}




}


func main() {
	head := make([]byte,4)
	openFile := fileOpen("\\goworkspace\\tmp\\dat")
	defer openFile.Close()
	_, e2 := openFile.Read(head)
	errorCheck(e2)
	typeCheck(head)



}

