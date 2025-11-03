#!/bin/bash

### Create the responce FIFO
rm -f response
mkfifo response

function handleRequest() {
    # 1) Process the request
    # 2) Route request to the correct handler
    # 3) Build a response based on the request
    # 4) Send the response to the named pipe (FIFO)

	while read line; do
    		trline=$(echo "$line" | tr -d '[\r\n]')

    		if [ -z "$trline" ]; then
      			break
    		fi

		HEADLINE_REGEX='(.*?)\s(.*?)\sHTTP.*?'

		[[ "$trline" =~ $HEADLINE_REGEX ]] &&
		REQUEST=$(echo $trline | sed -E "s/$HEADLINE_REGEX/\1 \2/")

  	done

	case "$REQUEST" in                                                                   
		"GET /") RESPONSE="HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n<h1>PONG</h1>" 
				;;
		*) 
			RESPONSE="HTTP/1.1 404 Not Found\r\nContent-Type: text/html\r\n\r\n<h1>Not Found</h1>"
				;;
 	esac

	echo -e $RESPONSE > response
}

echo 'Listening on 3000...'

while true; do
	cat response | nc -l 3000 | handleRequest
done
