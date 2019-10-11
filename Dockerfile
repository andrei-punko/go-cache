# iron/go is the alpine image with only ca-certificates added
FROM iron/go

WORKDIR /app

# Now just add the binary
COPY ./out/linux-amd64/web-cache /app/

ENTRYPOINT ["./web-cache"]
