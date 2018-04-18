# Picomposer

Generate all variants from png layers.

### Why?

Sometimes you need to generate a lot of different pictures from few parts.
Picomposer combines layers of transparent png images into all possible variations

### How to

Just run 
(default folders are `./pngs` for input folder and `./output` for output folder)
```
./picomposer
```
And then wait for all variants stored in output folder

### Flags
`--input-folder` - specify input folder with given structure
`--output-folder` - specify output folder for generated pictures

### Folder structure

You need folder structure for `input folder`:
```
./input-folder/
  ../layer001/
    ../pic001.png
    ../pic002.png
    ../pic003.png
  ../layer002/
    ../pic004.png
    ../pic005.png
```