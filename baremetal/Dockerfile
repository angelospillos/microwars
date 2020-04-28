# ------------------------------------------------------------------------------
# Build Stage
# ------------------------------------------------------------------------------
FROM rust:latest as cargo-build
WORKDIR /app
ADD . /app
RUN cargo build --release

# ------------------------------------------------------------------------------
# Production Stage
# ------------------------------------------------------------------------------
FROM gcr.io/distroless/cc
COPY --from=cargo-build /app/target/release/baremetal /
CMD ["./baremetal"]