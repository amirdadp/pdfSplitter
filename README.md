# pdfSplitter
A small bash script to split pdfs using GhostScript with a more friendly experience.
(I made this program to help me spliting assignment questions to submit to Crowdmark platform.)


# Requirements
* Ghostscript


# How to use

The script takes 2+ arguments. The first argument is the name of the pdf file that you want to split. (This can be a path)

The rest of the arguments are the page numbers of the last page of each section. For example;

```
./pdfSplitter.sh FILENAME.pdf 2 4 5
```
 would split `FILENAME.pdf` into three files: 
 
 1. Page 1 to Page 2
 2. Page 3 to Page 4
 3. Page 5 to Page 5
