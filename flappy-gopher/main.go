package main

import (
	"fmt"
	"os"

	"time"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"

	img "github.com/veandco/go-sdl2/sdl_image"
)

func run() error {
	// Initialize SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Could not initialize SDL: %v", err)
	}
	// Defer closing of the program
	defer sdl.Quit()

	// Initialize sdl_ttf
	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("Could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	// Create the window and renderer
	w, r, err := sdl.CreateWindowAndRenderer(1280, 720, sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("Could not create window: %v", err)
	}
	defer w.Destroy()

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("Could not draw title: %v", err)
	}

	time.Sleep(5 * time.Second)

	if err := drawBackground(r); err != nil {
		return fmt.Errorf("Could not draw background: %v", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()
	f, err := ttf.OpenFont("res/fonts/Barrio-Regular.ttf", 20)
	if err != nil {
		return fmt.Errorf("Could not load font: %v", err)
	}
	defer f.Close()
	c := sdl.Color{
		R: 201,
		G: 100,
		B: 0,
		A: 255}

	s, err := f.RenderUTF8_Solid("Flappy Gopher", c)
	if err != nil {
		return fmt.Errorf("Could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Could not copy texture: %v", err)
	}
	r.Present()

	return nil
}

func drawBackground(r *sdl.Renderer) error {
	r.Clear()

	// Load the texture into the renderer
	t, err := img.LoadTexture(r, "res/img/background.png")
	if err != nil {
		return fmt.Errorf("Could not load background: %v", err)
	}

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Could not copy background: %v", err)
	}
	defer t.Destroy()

	r.Present()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
