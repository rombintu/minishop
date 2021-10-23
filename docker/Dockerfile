FROM alpine:1.14
RUN apk update && apk add go vim git
RUN git clone https://github.com/rombintu/minishop.git
WORKDIR /minishop
COPY ./config/config.toml /minishop/config/config.toml
RUN make build
CMD ["./main"]