cli:
	go build -mod vendor -o bin/twilio-handler cmd/twilio-handler/main.go
	go build -mod vendor -o bin/flatten-definition cmd/flatten-definition/main.go

lambda:
	@make lambda-twilio-handler

lambda-twilio-handler:
	if test -f main; then rm -f main; fi
	if test -f twilio-handler.zip; then rm -f twilio-handler.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/twilio-handler/main.go
	zip twilio-handler.zip main
	rm -f main
