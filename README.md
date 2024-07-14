# Tiktoken Token File Cache

Pre-downloaded cache of Tiktoken files to speed up application startup,
and also because I was slightly annoyed when the "Offline BPE Loader" didn't
work on the first try.

## Use

1. Pull this image into your build process and copy `/var/cache/tiktoken/*` to your image. 
2. Set the `TIKTOKEN_CACHE_DIR` environment variable.

## Example Dockerfile

```Dockerfile
FROM tylergannon/token-cache AS tiktoken

FROM golang:alpine

ENV TIKTOKEN_CACHE_DIR=/var/cache/tiktoken
COPY --from=tiktoken ${TIKTOKEN_CACHE_DIR}/. ${TIKTOKEN_CACHE_DIR}
