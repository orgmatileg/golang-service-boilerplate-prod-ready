FROM golang:1.16.2-alpine3.13

ENV TZ=Asia/Jakarta

RUN apk update && apk add --no-cache git ca-certificates binutils libgcc libstdc++ libx11 glib libxrender libxext libintl ttf-freefont fontconfig wkhtmltopdf ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family tzdata

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN update-ca-certificates

# Create appuser.
RUN adduser -D -g '' app
WORKDIR /app/

# Copy depedencies
COPY go.mod go.sum ./

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

ENTRYPOINT ["/go/bin/app","serve-api"]