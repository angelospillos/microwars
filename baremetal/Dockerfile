# ------------------------------------------------------------------------------
# Build Stage
# ------------------------------------------------------------------------------
FROM rust:latest as cargo-build
RUN apt-get update
RUN apt-get install musl-tools -y
RUN rustup target add x86_64-unknown-linux-musl
WORKDIR /usr/src/baremetal
COPY . .
RUN RUSTFLAGS=-Clinker=musl-gcc cargo build --release --target=x86_64-unknown-linux-musl
RUN rm -f target/x86_64-unknown-linux-musl/release/deps/baremetal*
COPY . .
RUN RUSTFLAGS=-Clinker=musl-gcc cargo build --release --target=x86_64-unknown-linux-musl

# ------------------------------------------------------------------------------
# Production Stage
# ------------------------------------------------------------------------------
FROM alpine:latest
RUN addgroup -g 1000 baremetal
RUN adduser -D -s /bin/sh -u 1000 -G baremetal baremetal
WORKDIR /home/baremetal/bin/
COPY --from=cargo-build /usr/src/app/target/x86_64-unknown-linux-musl/release/baremetal .
RUN chown baremetal:baremetal baremetal
USER baremetal
CMD ["./baremetal"]
