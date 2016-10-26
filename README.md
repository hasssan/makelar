# Makelar

Little program to receive webhook than trigger hugo build

Usage:

```
MLR_PORT=8080 \
MLR_URL=/webhook \
MLR_HUGO_SITE=~/path/to/hugo/site \
MLR_HUGO_BIN=/path/to/bin/hugo \
MLR_OUTPUT_DIR=/path/to/public_html \
./makelar
```