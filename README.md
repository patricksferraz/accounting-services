<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
***
***
***
*** To avoid retyping too much info. Do a search and replace for the following:
*** github_username, repo_name, twitter_handle, email, project_title, project_description
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/c4ut/accounting-services">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Accounting Services</h3>

  <p align="center">
    Repository containing microservices that contemplate the infrastructure of a system for an accounting firm
    <br />
    <a href="https://github.com/c4ut/accounting-services"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/c4ut/accounting-services">View Demo</a>
    ·
    <a href="https://github.com/c4ut/accounting-services/issues">Report Bug</a>
    ·
    <a href="https://github.com/c4ut/accounting-services/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <!-- <li><a href="#usage">Usage</a></li> -->
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <!-- <li><a href="#license">License</a></li> -->
    <li><a href="#contact">Contact</a></li>
    <!-- <li><a href="#acknowledgements">Acknowledgements</a></li> -->
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

Project developed using the DDD (Domain Driven Design) design pattern containing microservices that meet the needs of an accounting firm, currently with the following services:

- [authentication service](https://github.com/c4ut/accounting-services/service/auth)
- [point registration service](https://github.com/c4ut/accounting-services/service/time-record)
<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->
<!--
Here's a blank template to get started:
**To avoid retyping too much info. Do a search and replace with your text editor for the following:**
`github_username`, `repo_name`, `twitter_handle`, `email`, `project_title`, `project_description` -->

### Built With

- [Go Lang](https://golang.org/)
- List all: `go list -m all`

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Hiring a kubernetes cluster:
  - [AWS](https://aws.amazon.com/pt/eks/?whats-new-cards.sort-by=item.additionalFields.postDateTime&whats-new-cards.sort-order=desc&eks-blogs.sort-by=item.additionalFields.createdDate&eks-blogs.sort-order=desc)
  - [Azure](https://azure.microsoft.com/pt-br/services/kubernetes-service/)
  - [GCP](https://cloud.google.com/kubernetes-engine)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- Create a secret for github docker registry

  ```sh
  kubectl create secret docker-registry regcred \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL
  ```

- Create a secret with env credentials

  ```sh
  # credentials for auth-service
  # file: credentials
  KEYCLOAK_REALM=keycloak_realm
  KEYCLOAK_CLIENT_ID=keycloak_client_id
  KEYCLOAK_CLIENT_SECRET=keycloak_client_secret
  KEYCLOAK_AUDIENCE=account
  ```

  `kubectl create secret generic auth-secret --from-env-file ./credentials`

  ```sh
  # credentials for time-record-service
  # file: credentials
  DB_URI=mongodb://user:password@mongo
  DB_NAME=time_record_service
  ```

  `kubectl create secret generic time-record-secret --from-env-file ./credentials`

### Installation

- `kubectl apply -f ./k8s`

<!-- USAGE EXAMPLES -->
<!-- ## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_ -->

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/c4ut/accounting-services/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

__Prerequisites__:

- Golang

  ```sh
  wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz
  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
  export PATH=$PATH:/usr/local/go/bin
  ```

- Docker and docker-compose

  ```sh
  sudo apt-get install docker docker-compose docker.io -y
  ```

- Environment

  ```sh
  # .env
  AUTH_API_PORT=50051

  AUTHDB_DB=keycloak
  AUTHDB_USERNAME=keycloak
  AUTHDB_PASSWORD=password

  KEYCLOAK_USERNAME=admin
  KEYCLOAK_PASSWORD=Pa55w0rd
  KEYCLOAK_BASE_PATH=http://keycloak:8080
  KEYCLOAK_REALM=keycloak_realm
  KEYCLOAK_CLIENT_ID=keycloak_client_id
  KEYCLOAK_CLIENT_SECRET=keycloak_client_secret

  TIME_RECORD_API_PORT=50052

  MONGODB_USERNAME=admin
  MONGODB_PASSWORD=admin123
  DB_URI=mongodb://admin:admin123@trdb:27017
  DB_NAME=time_record_service
  DB_PORT=27018
  DB_TEST=true # for unit and integration tests
  DB_MIGRATE=true # to migrate "up" at database startup

  AUTH_SERVICE_ADDR=auth-service:50051
  ```

__Installation__:

1. Clone the repo

   ```sh
   git clone https://github.com/c4ut/accounting-services.git
   ```

2. Run

   ```sh
   docker-compose up -d
   ```

3. Test

   ```sh
   go test -v -coverprofile cover.out ./...
   go tool cover -html=cover.out -o cover.html
   ```

__Local kubernetes__:

1. Install [Kind](https://kind.sigs.k8s.io/) or similar
2. Follow the steps of [Getting Started](#getting-started)
    - For the local keycloak, run:

      `kubectl create -f https://raw.githubusercontent.com/keycloak/keycloak-quickstarts/latest/kubernetes-examples/keycloak.yaml`

    - For the local mongodb, do:

      ```sh
      # create: mongo.yml
      apiVersion: apps/v1
      kind: StatefulSet
      metadata:
        name: mongo
      spec:
        serviceName: mongo
        replicas: 1
        selector:
          matchLabels:
            app: mongo
        template:
          metadata:
            labels:
              app: mongo
              selector: mongo
          spec:
            containers:
            - name: mongo
              image: mongo:4.4
              ports:
                - containerPort: 27017
              env:
                - name: MONGO_INITDB_ROOT_USERNAME
                  value: admin
                - name: MONGO_INITDB_ROOT_PASSWORD
                  value: password
            nodeSelector:
              kubernetes.io/hostname: accounting-control-plane
      ---
      apiVersion: v1
      kind: Service
      metadata:
        name: mongo
        labels:
          app: mongo
      spec:
        # clusterIP: None
        selector:
          app: mongo
        ports:
        - port: 27017
          targetPort: 27017
      ```

      run:

      `kubectl apply -f mongo.yaml`
<!-- LICENSE -->
<!-- ## License -->

<!-- Distributed under the MIT License. See `LICENSE` for more information. -->

<!-- CONTACT -->
## Contact

Patrick Ferraz - [@patricksferraz](https://twitter.com/patricksferraz) - comercial@coding4u.com.br

Project Link: [https://github.com/c4ut/accounting-services](https://github.com/c4ut/accounting-services)

<!-- ACKNOWLEDGEMENTS -->
<!-- ## Acknowledgements

* []()
* []()
* []() -->

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/44science/theplace?style=for-the-badge
[contributors-url]: https://github.com/c4ut/accounting-services/repo/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/44science/theplace?style=for-the-badge
[forks-url]: https://github.com/c4ut/accounting-services/repo/network/members
[stars-shield]: https://img.shields.io/github/stars/44science/theplace?style=for-the-badge
[stars-url]: https://github.com/c4ut/accounting-services/repo/stargazers
[issues-shield]: https://img.shields.io/github/issues/44science/theplace?style=for-the-badge
[issues-url]: https://github.com/c4ut/accounting-services/repo/issues
[license-shield]: https://img.shields.io/github/license/44science/theplace?style=for-the-badge
[license-url]: https://github.com/c4ut/accounting-services/repo/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/patricksferraz
