# ------------------------------------------------------------------------------
# Build Stage
# ------------------------------------------------------------------------------
FROM rust:latest as cargo-build
RUN apt-get update
RUN apt-get install musl-tools -y
RUN rustup target add x86_64-unknown-linux-musl
WORKDIR /usr/src/app
COPY . .
RUN RUSTFLAGS=-Clinker=musl-gcc cargo build --release --target=x86_64-unknown-linux-musl
RUN rm -f target/x86_64-unknown-linux-musl/release/deps/app*
COPY . .
RUN RUSTFLAGS=-Clinker=musl-gcc cargo build --release --target=x86_64-unknown-linux-musl

# ------------------------------------------------------------------------------
# Production Stage
# ------------------------------------------------------------------------------
FROM alpine:latest
RUN addgroup -g 1000 app
RUN adduser -D -s /bin/sh -u 1000 -G app app
WORKDIR /home/app/bin/
COPY --from=cargo-build /usr/src/app/target/x86_64-unknown-linux-musl/release/app .
RUN chown app:app app
USER app
CMD ["./app"]
