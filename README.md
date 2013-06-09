## Procedural Webserver

A webserver written in Google Go which serves up procedurally generated content.

I built this for a webcrawler which I am working on--I did not want to test 
it out by hamming the same set of websites over and over (as I am sure the
owners would not appreciate that) so I instead wrote this webserver to serve 
up content which is generated procedurally.  That way, I can hammer it with my 
web crawler yet still get unique pages generated with little memory usage.

### Usage

To use this in your app, you'd want Go code that looks like this:

    import server "github.com/dmuth/procedural-webserver"

    func main() {
        server_object := server.NewServer(port, NumLinksMin, NumLinksMax, NumImagesMin, NumImagesMax, Seed)
        go server_object.Start() // Start the webserver
        // do stuff
        server_object.Stop() // Stop the webserver
    }

- NumLinksMin: Min number of links. (actual number of created images is somewhere between min and max)
- NumLinksMax: Max number of links listed on each generated page. (links bring up more procedurally generated content)
- NumImagesMin: Min number of images. (actual number of created images is somewhere between min and max)
- NumImagesMax: Max number of images listed on each generated page. (they're not really images)
- Seed: Our base seed. If not specified, a default is used.  If specified, this change s the content of 
all pages. If you understand how Minecraft world seeds work, this is the same concept.
    
Once the server is running, do `curl localhost` to bring up the main page and go from there.

### Query parameters

The following query parameters will have effects on the page that is generated:

- code=num - This will force the websever to return a specific HTTP code.  
    `code=404` will return a 404, for example.
- delay=num How long to delay loading a page for.  Useful for simulating network lag
    - Example times: 
        - 100ms
        - 2s
        - 5m
        - 1h (Yes, this is an hour.  Google Go doesn't care. But Curl might!)
 
### TODO list for The Future
- Optimizations under heavy load
- Generate a random title for each page
- Add support for making page links include a 404 code a certain percentage of the time
- Add support for making page links ionclude a 3xx code a ceratin percentage of the time

### Comments and complaints!

Send them to me, Douglas Muth: http://www.dmuth.org/contact

Filing a bug here is fine, too!

Either way, I'd love to know if you're using this webserver, and find it helpful.



