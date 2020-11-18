#! /bin/bash

docker pull blang/latex:ctanfull

if [ -f "latexdockercmd.sh" ]; then
    echo "Script already exists locally."    
else
    echo "Downloading script."
    wget https://raw.githubusercontent.com/blang/latex-docker/master/latexdockercmd.sh
fi
chmod +x latexdockercmd.sh

mkdir pdf
./latexdockercmd.sh pdflatex -interaction nonstopmode -output-directory ./pdf main.tex

mv ./pdf/main.pdf Zachary-Rohrbach-Resume.pdf
rm -rf pdf

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    xdg-open Zachary-Rohrbach-Resume.pdf
elif [[ "$OSTYPE" == "darwin"* ]]; then
    open Zachary-Rohrbach-Resume.pdf
fi