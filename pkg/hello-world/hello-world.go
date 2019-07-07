package helloworld

import(
	"github.com/common-nighthawk/go-figure"
)

func HelloWorld() {
	myFigure := figure.NewFigure("Hello World", "", true)
  myFigure.Print()
}
