# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy source
COPY main.go .

# Build statically-linked binary
RUN CGO_ENABLED=0 go build -o server main.go


# ---------- Runtime stage ----------
FROM scratch

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Copy servers.json if you want it baked into the image
# (can also mount at runtime instead)
# COPY servers.json .

# Expose port
EXPOSE 8080

# Run
ENTRYPOINT ["/app/server"]

