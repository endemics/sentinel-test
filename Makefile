.PHONY: all clean

all:
	docker build -t sentinel-test .
	docker run --rm -ti sentinel-test

clean:
	docker rmi -f sentinel-test
