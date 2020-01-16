# iron/go is the alpine image with only ca-certificates added
FROM iron/go

WORKDIR /app

# Now just add the binary
COPY ./.gogradle/linux_amd64_go-cache /app/web-cache

ENTRYPOINT ["./web-cache"]
