FROM rust:1.79.0

WORKDIR /usr/src/app

COPY . .

RUN cargo build --release

EXPOSE 8080

CMD ["./target/release/k8s-simple-rust-project"]