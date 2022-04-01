serve:
	go build ./server
	./server.exe


app:
	go run ./gui

clean:
	del  .\\bucket\*
