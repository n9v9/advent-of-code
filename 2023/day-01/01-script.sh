#!/usr/bin/env sh

rg '(?:^\D*?(\d).*(\d))|^\D*((\d))\D*' -o -r '$1 $2 $3 $4' input.txt |
	rg '^\s*(\d) (\d)\s*' -r '$1$2' |
	paste -d + -s |
	bc
