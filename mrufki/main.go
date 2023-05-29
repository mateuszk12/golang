package main

import (
	"flag"
	"fmt"
	// "fmt"
	"math"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

type Point struct {
	x int32
	y int32
}
type Mrufka struct {
	loc   Point
	dir   Point
	prev  Point
	tail  int
	color uint32
	path int
	// status int
	//free - 0
	//carry - 1
}

func random(upper, lower int32) (random int32) {
	rng := upper - lower
	random = int32(rand.Intn(int(rng))) + lower
	return
}
func rotate(x int32, y int32, angle int) (xx, yy int32) {
	yy = int32(math.Round(math.Sin(float64(angle) * (math.Pi / 180))))
	xx = int32(math.Round(math.Cos(float64(angle) * (math.Pi / 180))))
	return
}
func rotateAcc(x float32, y float32, angle float32) (xx, yy float32) {
	yy = float32(math.Sin(float64(angle) * (math.Pi / 180)))
	xx = float32(math.Cos(float64(angle) * (math.Pi / 180)))
	return
}

// func (m Mrufka) (mrufka *Mrufka,w int32,h int32)
func move(mrufka *Mrufka, w int32, h int32) {
	lx := &mrufka.loc.x
	ly := &mrufka.loc.y
	dy := &mrufka.dir.y
	dx := &mrufka.dir.x
	px := &mrufka.prev.x
	py := &mrufka.prev.y
	path := &mrufka.path
	*px = *lx
	*py = *ly
	if (random(1000,-1000) % 20 == 0 && *path !=0) {
		*dx, *dy = *dx+random(2,-1),*dy+random(2,-1)
	}
	// if (*path%40==0 && *path !=0 && random(1000,-1000) == 50){
	// 	*dx, *dy = rotate(*dx, *dy, 180)
	// }
	*lx = (*lx + *dx) +random(2,-1)
	*ly = (*ly + *dy) +random(2,-1)
	// kontakt ze ścianą
	if *ly > h-5 || *ly < 1 {
		*ly = *py
		_, *dy = rotate(*dx, *dy, rand.Intn(180))
	}
	if *lx < 1 || *lx > w-5 {
		*lx = *px
		*dx, _ = rotate(*dx, *dy, rand.Intn(180))
	}
	*path += 1
}
func Sim(antNumber int, window *sdl.Window,framerate uint64) {
	var mrufki []Mrufka
	w, h := window.GetSize()
	fmt.Println(w,h)
	for i := 0; i < int(antNumber); i++ {
		//mrufki = append(mrufki, Mrufka{Point{w/2,h/2}, Point{random(2, -1), random(2, -1)}, Point{100, 100},rand.Intn(16777216), rand.Intn(16777216),0})
		mrufki = append(mrufki, Mrufka{Point{w/2,h/2}, Point{random(2, -1), random(2, -1)}, Point{100, 100}, 16777215,16777216,0})
	}
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0xffffff)
	window.UpdateSurface()

	running := true
	for running {
		var start uint64 = sdl.GetTicks64()
		for i := range mrufki {
			move(&mrufki[i], w, h)
			rect := sdl.Rect{X: mrufki[i].loc.x, Y: mrufki[i].loc.y, W: 5, H: 5}
			pRect := sdl.Rect{X: mrufki[i].prev.x, Y: mrufki[i].prev.y, W: 5, H: 5}

			surface.FillRect(&pRect, uint32(mrufki[i].tail))
			surface.FillRect(&rect, mrufki[i].color)
			
		}
		window.UpdateSurface()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("quit")
				running = false
				break
			}
		}
		var end uint64 = sdl.GetTicks64()
		if end-start < 1000/framerate{
			sdl.Delay(uint32(1000/framerate-(end-start)))
			
		} 
		
	}
}
func main() {
	maxx, maxy := flag.Int("maxx", 1000, ""), flag.Int("maxy", 1000, "")
	numOfAnts := flag.Int("n",100,"number of ants")
	fps := flag.Uint64("f",60,"framerate")
	flag.Parse()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	window, err := sdl.CreateWindow("window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(*maxx), int32(*maxy), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	Sim(*numOfAnts, window,*fps)
}
