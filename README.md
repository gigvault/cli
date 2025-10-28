# GigVault CLI

Command-line interface for managing GigVault Certificate Authority operations.

## Installation

```bash
go install github.com/gigvault/cli/cmd/gigvault@latest
```

Or build locally:

```bash
make build
sudo cp bin/gigvault /usr/local/bin/
```

## Usage

### Certificate Operations

```bash
# Create a self-signed certificate
gigvault cert create example.com -o ./certs

# List certificates
gigvault cert list

# Revoke a certificate
gigvault cert revoke <serial>
```

### CSR Operations

```bash
# Create a new CSR
gigvault csr create example.com -o ./csrs

# Submit CSR to RA
gigvault csr submit ./csrs/example.com.csr
```

### Key Management

```bash
# Generate a new key pair
gigvault key generate mykey -c P-256

# Generate P-384 key
gigvault key generate mykey -c P-384
```

### Configuration

```bash
# Show current configuration
gigvault config show

# Set CA endpoint
gigvault config set ca-url https://ca.example.com
```

## Development

```bash
# Build
make build

# Run tests
make test

# Install locally
make install
```

## License

Copyright © 2025 GigVault

