# Makelar

![Unmaintained](https://img.shields.io/badge/status-unmaintained-red.svg)

Little program to receive webhook then update repository, 
build [hugo](http://gohugo.io/) site to output directory.

This project just my naive implementation for learning golang.

Usage:

```
MLR_PORT=8080 \
MLR_URL=/webhook \
MLR_HUGO_SITE=~/path/to/hugo/site \
MLR_HUGO_BIN=/path/to/bin/hugo \
MLR_OUTPUT_DIR=/path/to/public_html \
./makelar
```