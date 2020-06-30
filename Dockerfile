FROM golang:1.12.4

ENV CGO_ENABLED 0

WORKDIR $GOPATH/src/github.com/personal-work/space_exploration

RUN addgroup --system projects && adduser --system projects --ingroup projects

RUN chown -R projects:projects $GOPATH/src/github.com/personal-work/space_exploration

USER projects

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

RUN go install -v ./...

# This container exposes port 9001 to the outside world
EXPOSE 9001

# Run the executable
CMD ["space_explore"]

