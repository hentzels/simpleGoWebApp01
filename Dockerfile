FROM golang:alpine
COPY wiki.go /
RUN go build -o /wiki /wiki.go

# Copy the exe into a smaller base image
FROM alpine
COPY --from=0 /wiki /wiki
CMD /wiki
