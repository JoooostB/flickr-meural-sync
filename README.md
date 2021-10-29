# Flickr to NETGEAR Meural sync
Synchronise a Flickr Album of choice with your [NETGEAR Meural](https://www.netgear.com/nl/home/digital-art-canvas/) using this tool. Written in Go.

## Prerequisites
* Meural account
* Flickr API credentials

## Setup
This tool makes use of the Flickr- and a (reverse-engineered) Meural API. In order for it to function properly you need to set a few environment variables:
| Name              | Note                                       |
|-------------------|--------------------------------------------|
| MEURAL_EMAIL      |                                            |
| MEURAL_PASSWORD   | Make sure to escape any special characters |
| MEURAL_GALLERY_ID | Gallery to upload Flickr photos to         |
| FLICKR_KEY        |                                            |
| FLICKR_SECRET     |                                            |
`.env.example` is available to use as a base.

## Usage
Easiest method for running `flickr-meural-sync` is with the help of Docker:
```bash
docker run -it --env-file=.env joooostb/flickr-meural-sync:latest
```

Of course you're also able to use a binary for your platform or build and run the binary yourself:
```bash
go build main.go
./main
```