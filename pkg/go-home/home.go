package gohome

import("github.com/common-nighthawk/go-figure")

func GoHome() {
	figure.NewFigure("Go Home!", "basic", true).Scroll(30000, 400, "right")
}
