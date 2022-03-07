package main

import "github.com/fogleman/gg"

func main() {
	dc := gg.NewContext(500, 500)
	dc.DrawCircle(250, 250, 25)
	dc.SetRGB(1, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}
