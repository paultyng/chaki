# Chaki

## Task Format

```yaml
dbConnections:
  sqlite:
    driver: sqlite3
    dataSource: ':memory:'
tasks:
  update-user-email:
    title: Update User's Email
    description: Update a user's email address.
    db:
      connection: sqlite
      sql: update users set email = :email where id = :id
    schema:
      properties:
        id:
          title: User ID
          type: integer
        email:
          title: Email
          pattern: '@'
          type: string
  update-order-status:
    title: Update Order Status
    description: Change an order's status.
    db:
      connection: sqlite
      sql: update orders set status = :status where number = :number
    schema:
      properties:
        number:
          title: Order Number
          type: string
          pattern: '[0-9]+'
        status:
          title: Status
          type: string
          enum:
            - shipped
            - delivered
            - cancelled
          default: shipped
```

## Setup

Generate a key for the JWT:

```bash
ssh-keygen -t rsa -b 4096 -f chaki.key
# Don't add passphrase
openssl rsa -in chaki.key -pubout -outform PEM -out chaki.key.pub
cat chaki.key
cat chaki.key.pub
```

## Running

```bash
go install
yarn build
chaki serve
```
