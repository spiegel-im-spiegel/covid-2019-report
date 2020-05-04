package imgutil

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func ConcatImageFiles(dstPath string, srcPaths ...string) error {
	srcImages := []image.Image{}
	width, height := 0, 0
	for _, path := range srcPaths {
		img, err := imageFrom(path)
		if err != nil {
			return err
		}
		rct := img.Bounds()
		width = max(width, rct.Dx())
		height += rct.Dy()
		srcImages = append(srcImages, img)
	}

	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))
	offset := 0
	for _, img := range srcImages {
		srcRect := img.Bounds()
		draw.Draw(
			dstImage,
			image.Rect(0, offset, srcRect.Dx(), offset+srcRect.Dy()),
			img,
			image.Point{0, 0},
			draw.Over,
		)
		offset += srcRect.Dy()
	}

	file, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := png.Encode(file, dstImage); err != nil {
		return err
	}
	return nil
}

func imageFrom(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
