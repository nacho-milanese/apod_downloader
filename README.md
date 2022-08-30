# apod_downloader
Downloads Astronomy Picture of the Day  from nasa's site


## The APOD site
Nasa's astronomy picture of the day is a site which shares a picture, as it says, per day.
Some times is a video or more than one picture, so, by now for these days this application will crash :D

## The URL
The apod url is https://apod.nasa.gov/
I recommend you visiting the site, as besides the picture (and it's credits) you will find an interesting explanation of it, and some times, links to the topic.

## The purpose
I made this just to practice a little bit with golang. I've coded this before for python in order to gather screensaver/desktop background good quality beautiful images. Is just an excersise as mentioned.
I suggest following my advice in "The URL".
All credits on the content you can download with this is mentioned in the page, I just provide a mean to keep an image that is automatically downloaded when loading the page, there are no hacks involved in this.

## Usage

compile the source with `go build apod.go` or just run it by `go run apod.go`
The application requires a subdir `image_download` located in the same place as the code/binary (it won't be created if not there as this is expected to be a trivial application)

**Enjoy!**