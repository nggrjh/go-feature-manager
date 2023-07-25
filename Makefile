test:
	# source: https://gist.github.com/gregohardy/d148db9e01aa721ea42daf4abbba725e
	echo "\033[34mRunning test...\033[39m"
	go generate ./...
	go test ./... -count=1