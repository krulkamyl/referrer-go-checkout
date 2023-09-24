
# Refferer Golang

This is a GoLang project for managing orders and calculating commissions for referrers and administrators. The project structure is designed to resemble that of Laravel. It utilizes technologies like REDIS, MySQL, and MailHog. You can easily set up the project using the following Docker Compose command:

```bash
docker-compose build && docker-compose up
```

# Prerequisites

Before you start, make sure you have the following prerequisites installed on your system:

* Docker
* Docker Compose


# Seed Data
To populate the database with initial data, use the following commands:
```bash
go run app/console/commands/populateUsers.go
go run app/console/commands/populateProducts.go
go run app/console/commands/populateOrders.go
```

# Update ranking
To update rankings, run the following command:
```bash
go run app/console/commands/updateRankings.go
```