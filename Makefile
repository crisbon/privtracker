build:
	CGO_ENABLED=1 go build -ldflags="-s -w -buildid=" -trimpath

install:
	cp -f privtracker /usr/local/bin/privtracker

deploy: build
	rsync -avzL --exclude '*.gz' docs privtracker privtracker:web/

test:
	go test -bench . -benchmem
