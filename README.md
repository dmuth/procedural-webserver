## procedural-webserver

A webserver written in Google Go which served up procedurally generated content.

I built this for a webcrawler which I am working on--I did not want to test 
it out by hamming the same set of websites over and over (as I am sure the
owners would not appreciate that) so I instead wrote this webserver to serve 
up content which is generated procedurally.  That way, I can hammer it with my 
web crawler yet still get unique pages generated with little memory usage.


### Syntax

`go run ./main.go --debug-levellevel --num-images-max n --num-images-min n --num-links-max n --num-links-min n --seed "seed string"
