#!/usr/bin/env sh

digit='\d|one|two|three|four|five|six|seven|eight|nine'

rg "(?:^\D*?($digit).*($digit))|^\D*(($digit))\D*" -o -r '$1 $2 $3 $4' input.txt |
	rg "^\s*($digit) ($digit)\s*" -r '$1$2' |
	sed 's#one#1#g' |
	sed 's#two#2#g' |
	sed 's#three#3#g' |
	sed 's#four#4#g' |
	sed 's#five#5#g' |
	sed 's#six#6#g' |
	sed 's#seven#7#g' |
	sed 's#eight#8#g' |
	sed 's#nine#9#g' |
	paste -d + -s |
	bc
