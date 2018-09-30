# queue
This is an integer queue! It's very simple and is a learning project for me. As far as I can tell, it's a dequeue using a doubly linked list.

A cutesy program that lets you play with this queue exists at https://github.com/chrobid/queueplay

To use queueplay to play with this queue, [install Go](https://golang.org/doc/install), then:

    cd ~/go/src/
    git clone https://github.com/chrobid/queueplay.git
    cd queueplay/
    go get
    go build
    ./queueplay

My main goals were:
* Implement a queue without using arrays as the underlying data structure
* Package it, put it on github, and use it in another program while adhering to Go idioms wrt packaging and distribution

I had never done any of these things before, and I learned so much in the process!
