.PHONY: all

all:
	docker build -t sentinel-test .
	docker run --rm -ti sentinel-test
