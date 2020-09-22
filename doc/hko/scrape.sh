#!/bin/bash

# https://www.hko.gov.hk/tc/gts/time/conversion1_text.htm

for y in {1901..2100}; do
	url=https://www.hko.gov.hk/tc/gts/time/calendar/text/files/T${y}c.txt
	echo $url
    curl -LO# $url
done
