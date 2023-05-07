FROM golang:1.18 as builder

COPY . /gohtmltable
WORKDIR /gohtmltable/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzz_gohtmltable.a && \
    clang -fsanitize=fuzzer fuzz_gohtmltable.a -o fuzz_gohtmltable.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /gohtmltable/mayhem/fuzz_gohtmltable.libfuzzer /