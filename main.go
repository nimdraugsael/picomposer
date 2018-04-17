package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"image"
	"github.com/fogleman/gg"
	"strings"
	"strconv"
	"flag"
	"os"
)

var (
	inputFolder    = flag.String("input-folder", "./pngs", "folder with all images for composition")
	outputFolder    = flag.String("output-folder", "./output", "output folder")
)


type Picomposer struct {
	assetsFolder string
	outputFolder string
	tuples [][]namedImage
	images map[int][]namedImage
}

type namedImage struct {
	img image.Image
	name string
}


func (p *Picomposer) buildTuples(n int, res []namedImage) {
	if p.images[n] == nil {
		p.tuples = append(p.tuples, res)
		return
	}
	for _, node := range p.images[n] {
		next := append(res, node)
		p.buildTuples(n + 1, next)
	}
}

func (p *Picomposer) loadImages() {
	files, err := ioutil.ReadDir(p.assetsFolder)
	if err != nil {
		log.Fatal(err)
	}

	for i, layerFolder := range files {
		if layerFolder.IsDir() {
			imgFiles, err := ioutil.ReadDir(p.assetsFolder + "/" + layerFolder.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, f := range imgFiles {
				img, err := gg.LoadPNG(p.assetsFolder + "/" + layerFolder.Name() + "/" + f.Name())
				if err != nil {
					log.Fatal(err)
				}
				name := strings.Split(f.Name(), ".")[0]
				p.images[i] = append(p.images[i], namedImage{img, name})
			}
		}
	}
}

func (p *Picomposer) generateImages() {
	fmt.Println("Image generation started to: ", p.outputFolder)
	if _, err := os.Stat(p.outputFolder); os.IsNotExist(err) {
		fmt.Println("Output folder doesn't exist, creating")
		os.Mkdir(p.outputFolder, os.ModePerm)
	}

	for i, t := range p.tuples {
		ctx := gg.NewContextForImage(t[0].img)
		name := "out_"
		for _, ni := range t {
			ctx.DrawImage(ni.img, 0, 0)
			name = name + "_" + ni.name
		}
		path := p.outputFolder + "/" + name + ".png"
		ctx.SavePNG(path)
		fmt.Println("Saved " + strconv.FormatInt(int64(i + 1), 10) + "/" + strconv.FormatInt(int64(len(p.tuples)), 10))
	}
}



func main() {
	flag.Parse()

	fmt.Println("Hello, picomposer")

	p := Picomposer{
		assetsFolder: *inputFolder,
		outputFolder: *outputFolder,
		tuples: make([][]namedImage, 0),
		images: make(map[int][]namedImage),
	}

	p.loadImages()

	p.buildTuples(0, make([]namedImage, 0))

	p.generateImages()

	fmt.Println(len(p.tuples))
}
