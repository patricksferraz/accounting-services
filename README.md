# Accounting Services

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/accounting-services)](https://goreportcard.com/report/github.com/patricksferraz/accounting-services)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/patricksferraz/accounting-services?status.svg)](https://godoc.org/github.com/patricksferraz/accounting-services)

A modern, scalable microservices-based accounting system built with Go, gRPC, and MongoDB. This project demonstrates best practices in microservices architecture, clean code, and cloud-native development.

## 🚀 Features

- **Microservices Architecture**: Modular design with separate services for authentication and time recording
- **gRPC Communication**: High-performance RPC framework for service-to-service communication
- **MongoDB Integration**: Flexible document database for data storage
- **Keycloak Authentication**: Enterprise-grade identity and access management
- **Elastic APM**: Application performance monitoring and tracing
- **Docker & Kubernetes**: Containerized deployment with orchestration support
- **Clean Architecture**: Well-structured codebase following domain-driven design principles

## 🏗️ Architecture

The project follows a clean architecture pattern with the following components:

```
.
├── client/                 # Client application
│   ├── application/       # Application services
│   ├── domain/           # Domain models and interfaces
│   ├── infrastructure/   # Infrastructure implementations
│   └── cmd/              # Command-line interface
├── service/              # Microservices
│   ├── auth/            # Authentication service
│   ├── time-record/     # Time recording service
│   └── common/          # Shared utilities
├── utils/               # Common utilities
└── k8s/                # Kubernetes configurations
```

## 🛠️ Technology Stack

- **Backend**: Go 1.16+
- **API**: gRPC
- **Database**: MongoDB
- **Authentication**: Keycloak
- **Monitoring**: Elastic APM
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Testing**: Go testing framework with testify

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- MongoDB
- Keycloak server

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/patricksferraz/accounting-services.git
   cd accounting-services
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the services using Docker Compose:
   ```bash
   docker-compose up -d
   ```

### Development

1. Start the development environment:
   ```bash
   make dev
   ```

2. Run tests:
   ```bash
   make test
   ```

3. Build the project:
   ```bash
   make build
   ```

## 📚 API Documentation

The project uses Protocol Buffers (protobuf) for API definitions. The service interfaces are defined in the following proto files:

- Authentication Service: `service/common/protofiles/auth.proto`
- Time Record Service: `service/common/protofiles/time_record.proto`

To generate the API documentation or client/server code from the protobuf definitions, you can use tools like [protoc](https://grpc.io/docs/protoc-installation/), [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go), or [protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Sferraz** - *Initial work* - [GitHub Profile](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [gRPC](https://grpc.io/)
- [MongoDB](https://www.mongodb.com/)
- [Keycloak](https://www.keycloak.org/)
- [Elastic APM](https://www.elastic.co/apm)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
