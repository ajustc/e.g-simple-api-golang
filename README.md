<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/6295/6295417.png" width="100" />
</p>
<p align="center">
    <h1 align="center">E.G-SIMPLE-API-GOLANG</h1>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/ajustc/e.g-simple-api-golang?style=flat&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/ajustc/e.g-simple-api-golang?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/ajustc/e.g-simple-api-golang?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/ajustc/e.g-simple-api-golang?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
		<em>Developed with the software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
	<img src="https://img.shields.io/badge/NOW-001211.svg?style=flat&logo=NOW&logoColor=white" alt="NOW">
</p>
<hr>

##  Quick Links

> - [ Repository Structure](#repository-structure)
> - [ Getting Started](#getting-started)
>   - [ Installation](#installation)

---

##  Repository Structure

```sh
└── e.g-simple-api-golang/
    ├── Dockerfile
    ├── README.md
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── pkg
        ├── api
        │   └── handlers.go
        ├── db
        │   └── connection.go
        ├── middleware
        │   └── json.go
        └── models
            └── article.go
```

---

##  Getting Started

***Requirements***

Ensure you have the following dependencies installed on your system:

* **Go**: `version x.y.z`

###  Installation

1. Clone the e.g-simple-api-golang repository:

```sh
git clone https://github.com/ajustc/e.g-simple-api-golang
```

2. Change to the project directory:

```sh
cd e.g-simple-api-golang
```

3. Install the dependencies:

```sh
docker compose up
```