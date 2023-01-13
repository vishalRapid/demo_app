FROM golang:1.10

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/vishalrana9915/demo_app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 9200

# Run the executable
RUN make dev