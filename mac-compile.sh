#! /bin/bash

mkdir pdf
pdflatex -halt-on-error -output-directory ./pdf main.tex
mv pdf/main.pdf Zachary-Rohrbach-Resume.pdf
rm -rf pdf
open Zachary-Rohrbach-Resume.pdf
