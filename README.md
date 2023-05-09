# pdfSplitter
A small CLI (<i>Originally a Bash script, now re-written in Go</i>) to split pdfs using GhostScript with a more friendly experience.
(I made this program to help me spliting assignment questions to submit to Crowdmark platform.)


# Requirements
* [Ghostscript](https://github.com/ArtifexSoftware/ghostpdl-downloads/releases/tag/gs9550): Any version between 9.10 and 9.55.0
* [Go](https://go.dev/dl/)

# How to use

1. Download the latest Release from the `Releases` tab. 
2. `unzip` the file and `cd` into the unzipped directory.
3. Build from the source code
```
$ go build main.go 
```
4. The binary is ready, enjoy! 

# Guide
- You can use the `--help` or `-h` flags to check the manual.

- Here is a small example: 
```
./pdfSplitter FILENAME.pdf 2 4 5
```
 would split `FILENAME.pdf` into three files: 
 
 1. Page 1 to Page 2
 2. Page 3 to Page 4
 3. Page 5 to Page 5
