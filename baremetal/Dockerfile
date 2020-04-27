# ------------------------------------------------------------------------------
# Build Stage
# ------------------------------------------------------------------------------
# Dockerfile for creating a statically-linked Rust application using docker's
# multi-stage build feature. This also leverages the docker build cache to avoid
# re-downloading dependencies if they have not changed.
FROM rust:latest as cargo-build
RUN apt-get update
RUN apt-get install musl-tools -y
WORKDIR /usr/src
# Download the target for static linking.
RUN rustup target add x86_64-unknown-linux-musl
# Create a dummy project and build the app's dependencies.
# If the Cargo.toml or Cargo.lock files have not changed,
# we can use the docker build cache and skip these (typically slow) steps.
RUN USER=root cargo new baremetal
WORKDIR /usr/src/baremetal
COPY Cargo.toml Cargo.lock ./
RUN cargo build --release
# Copy the source and build the application.
COPY src ./src
RUN cargo install --target x86_64-unknown-linux-musl --path .

# ------------------------------------------------------------------------------
# Production Stage
# ------------------------------------------------------------------------------
# Copy the statically-linked binary into a scratch container.
FROM scratch
COPY --from=cargo-build /usr/local/cargo/bin/baremetal .
USER 1000
CMD ["./baremetal"]