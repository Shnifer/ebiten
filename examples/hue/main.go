// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build example jsgo

package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	count        int
	gophersImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	count++

	if ebiten.IsRunningSlowly() {
		return nil
	}

	// Center the image on the screen.
	w, h := gophersImage.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth-w)/2, float64(screenHeight-h)/2)

	// Rotate the hue.
	op.ColorM.RotateHue(float64(count%360) * 2 * math.Pi / 360)

	screen.DrawImage(gophersImage, op)
	return nil
}

func main() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hue (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
