FROM golang:1.13.6-buster AS build-env

WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/perseal \
    -ldflags "-extldflags \"-fno-PIC -static \
      -lpthread -lstdc++\"" -buildmode pie -tags 'osusergo netgo static_build'

FROM scratch

COPY --from=build-env /go/bin/screenurls /go/bin/screenurls

ENTRYPOINT ["/go/bin/screenurls"]
