package main

import (
	"flag"
	"fmt"
	"github.com/Limard/gg"
	"testing"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

var (
	width  = flag.Int("w", 2100, "width")
	height = flag.Int("h", 2970, "height")
	dpi    = flag.Int("dpi", 300, "dpi")
)

func main() {
	flag.Parse()

	drawImageBenchmark()
}

func drawImageBenchmark() {
	img, e := gg.LoadImage(flag.Arg(0))
	if e != nil {
		panic(e)
	}

	dc := CreateDC()

	res := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dc = CreateDC()
		}
	})
	fmt.Println("CreateDC")
	fmt.Println("[RES]\t", res.String())
	fmt.Println("[MEM]\t", res.MemString())
	fmt.Println()

	res = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dc.DrawImage(img, 0, 0)
		}
	})
	fmt.Println("DrawImage")
	fmt.Println("[RES]\t", res.String())
	fmt.Println("[MEM]\t", res.MemString())
	fmt.Println()

	res = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dc.SavePNG("output.png")
		}
	})
	fmt.Println("SavePNG")
	fmt.Println("[RES]\t", res.String())
	fmt.Println("[MEM]\t", res.MemString())
	fmt.Println()
}

func CreateDC() (dc *gg.Context) {
	dc = gg.NewContext(Calc01mmToPt(*dpi, *width), Calc01mmToPt(*dpi, *height))

	// 填充白色
	dc.SetRGB255(0xFF, 0xFF, 0xFF)
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.Fill()

	dc.SetRGB255(0, 0, 0)
	return
}

func Calc01mmToPt(dpi, l int) int {
	return int(float64(l) / 10 / 25.4 * float64(dpi))
}
