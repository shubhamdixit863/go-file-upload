.PHONY: all build move dockerbuild start

all: dockerbuild


dockerbuild:
		@echo "building docker image"
		docker build --no-cache -t shubhamdixit863/gos3 .
		@echo "pushing docker image"
		docker push shubhamdixit863/gos3


