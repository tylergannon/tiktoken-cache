FROM golang:alpine as BASE
ENV TIKTOKEN_CACHE_DIR=/var/cache/tiktoken/
COPY . .
RUN go mod download && go run .

FROM scratch
ENV TIKTOKEN_CACHE_DIR=/var/cache/tiktoken/
COPY --from=BASE ${TIKTOKEN_CACHE_DIR}. ${TIKTOKEN_CACHE_DIR}
