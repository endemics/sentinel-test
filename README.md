This repository aims at building a test plugin for [Hashicorp sentinel](https://docs.hashicorp.com/sentinel/intro/) using the informations from the [Developing an Import Plugin](https://docs.hashicorp.com/sentinel/extending/dev) page.

## Running the plugin test

All you need is an internet connection and a machine with a copy of the repository as well as the commands `make` and `docker`.

You can then build the test image and run the test of the plugin with:
```
make all
```

You can then clean the produced docker image with: `make clean`

## Licence

There is no licence on the Hashicorp website so there won't be any licence for this repository either.

