FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates

RUN mkdir /pb
COPY ./eclosix /pb/

# uncomment to copy the local pb_migrations dir into the image
COPY ./migrations /pb/pb_migrations

# uncomment to copy the local pb_hooks dir into the image
# COPY ./pb_hooks /pb/pb_hooks

EXPOSE 8080

ENV ECLOSIX_ENV=production

# start PocketBase
CMD ["/pb/eclosix", "serve", "--http=0.0.0.0:8080"]