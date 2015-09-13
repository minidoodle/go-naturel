// Naturel scans images for pixels that are very simular to
// the color of skin
package naturel

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// Based on code from https://golang.org/pkg/image for reading
// image pixels and code from a pornscanner withen in python using
// the Python Image Library (PIL)

// Checks given file name for the skin ratio in the image
// returning if we think it's porn and what our skinRaito
// value came to
func IsPorn(imgName string) (isPorn bool, skinRatio float64, err error) {
	// Read in our Image
	reader, err := os.Open(imgName)
	if err != nil {
		return
	}
	defer reader.Close()

	// Decode the image
	m, _, err := image.Decode(reader)
	if err != nil {
		return
	}
	bounds := m.Bounds()

	var skin int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			// Look for a particular red color and variations of blue and green could reperesnt a skin color
			if r > 60 && float64(g) < (float64(r)*0.85) && float64(b) < (float64(r)*0.7) && float64(g) > (float64(r)*0.4) && float64(b) > (float64(r)*0.2) {
				skin++
			}
		}
	}

	size := bounds.Size()
	skinRatio = float64(skin) / float64(size.Y*size.X) * 100
	if skinRatio > 30 {
		isPorn = true
	}
	return
}
