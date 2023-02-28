# GraphQLServerTemplate

## Outline

Go 언어 기반의 GraphQL 서버 템플릿 프로젝트 아래와 같은 기능들이 템플릿으로 구현되어있다.
- MVC 기반의 템플릿 구조
  - M = Repository, V = Resolver, C = Service
- GraphQL Server
- Database
- Logger
- Env config
- Dataloader (cache)

## Installation

이 프로젝트를 빌드하기 위해서는 GO가 설치되어 있어야 한다.

이 프로젝트를 빌드하고, 실행하려거든 아래 커맨드를 수행하면 된다.

```bash
# Set up local database
make test_setup

# Excute
make start

# Test
make test
```

pkg 디렉터리 하위의 graph 디렉터리에 정의된 GraphQL 스키마를 변경할 경우 아래 커맨드를 통해 generate 한다.

```bash
# gqlgen
make gen (or make generate)
```
