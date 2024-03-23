## CHESZ
WIP Documentation

## Scaffolding
```
.
├── api
│   └── v1
├── cmd
│   ├── migration
│   ├── server
│   │   └── main.go
│   └── task
├── config
├── deploy
├── docs
├── internal
│   ├── handler
│   ├── middleware
│   ├── model
│   ├── repository
│   ├── server
│   └── service
├── pkg
├── scripts
├── test
│   ├── mocks
│   └── server
├── web
├── Makefile
├── go.mod
└── go.sum
```

`cmd:` Entry point of the application, containing different subcommands.
`config:` Configuration files.
`deploy:` Files related to deployment, such as Dockerfile and docker-compose.yml.
`internal:` Main code of the application
`pkg:` Common code, including configuration, logging, and HTTP.
`scripts:` Script files for deployment and other automation tasks.
`storage:` Storage files, such as log files.
`test:` Test code.
`web:` Front-end code.