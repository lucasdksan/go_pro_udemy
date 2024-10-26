## ACID

ACID é um conjunto de propriedades que garantem a confiabilidade e a consistência das operações realizadas em sistemas de banco de dados, especialmente em transações. A sigla representa quatro características principais: Atomicidade, Consistência, Isolamento e Durabilidade.

* **Atomicidade (A):** Garante que todas as operações de uma transação sejam tratadas como uma unidade indivisível. Se qualquer parte da transação falhar, nenhuma das operações será concluída. Isso significa que uma transação ocorre completamente ou não ocorre, mantendo o banco de dados em um estado consistente.

* **Consistência (C):** Assegura que uma transação leva o banco de dados de um estado válido para outro estado válido, cumprindo todas as regras de integridade e restrições do banco de dados. Isso significa que, após uma transação, os dados devem estar em um estado consistente, seguindo todas as regras definidas, como chaves estrangeiras e restrições.

* **Isolamento (I):** Define que as operações de diferentes transações não interferem umas nas outras. Mesmo se várias transações ocorrerem simultaneamente, elas não devem afetar os resultados entre si. A implementação de níveis de isolamento ajuda a evitar problemas como leituras sujas e atualizações perdidas.

* **Durabilidade (D):** Garante que, uma vez que uma transação é concluída, as alterações são permanentemente salvas no banco de dados, mesmo em caso de falhas no sistema. Isso é geralmente assegurado por meio de logs de transação ou backups, permitindo que o banco recupere o estado final após uma falha.

## Exemplo

Configuração:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    dsn := "postgres://postgres:docker@localhost:5432/testdb"
    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer pool.Close()

    fmt.Println("Database connected.")
    // Aqui você pode começar a criar funções para transações ACID
}
```

* **Atomicidade**

```go
func performAtomicTransaction(pool *pgxpool.Pool) error {
    ctx := context.Background()

    // Iniciar uma transação
    tx, err := pool.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %v", err)
    }
    defer tx.Rollback(ctx) // Garantia de rollback se algo falhar

    // Inserir novo usuário
    _, err = tx.Exec(ctx, "INSERT INTO users (id, name, balance) VALUES ($1, $2, $3)", 1, "John Doe", 1000)
    if err != nil {
        return fmt.Errorf("failed to insert user: %v", err)
    }

    // Atualizar saldo
    _, err = tx.Exec(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", 100, 1)
    if err != nil {
        return fmt.Errorf("failed to update balance: %v", err)
    }

    // Commit se tudo ocorrer bem
    err = tx.Commit(ctx)
    if err != nil {
        return fmt.Errorf("failed to commit transaction: %v", err)
    }
    return nil
}
```

**Consistência**

```go
func enforceConsistency(pool *pgxpool.Pool) error {
    ctx := context.Background()
    tx, err := pool.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    _, err = tx.Exec(ctx, "INSERT INTO accounts (id, balance) VALUES ($1, $2)", 1, 500)
    if err != nil {
        return fmt.Errorf("insert failed: %v", err)
    }

    // Tentativa de retirada que causaria saldo negativo, violando a constraint
    _, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", 600, 1)
    if err != nil {
        fmt.Println("Transaction failed due to constraint violation")
        return err
    }

    return tx.Commit(ctx)
}
```

**Isolamento**

```go
func performIsolatedTransaction(pool *pgxpool.Pool) error {
    ctx := context.Background()
    tx, err := pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %v", err)
    }
    defer tx.Rollback(ctx)

    // Ler o saldo
    var balance int
    err = tx.QueryRow(ctx, "SELECT balance FROM accounts WHERE id = $1", 1).Scan(&balance)
    if err != nil {
        return fmt.Errorf("failed to read balance: %v", err)
    }

    // Atualizar saldo
    _, err = tx.Exec(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", balance-100, 1)
    if err != nil {
        return fmt.Errorf("failed to update balance: %v", err)
    }

    return tx.Commit(ctx)
}
```

**Durabilidade**

```go
func ensureDurability(pool *pgxpool.Pool) error {
    ctx := context.Background()
    tx, err := pool.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    _, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", 500, 1)
    if err != nil {
        return fmt.Errorf("failed to update balance: %v", err)
    }

    // Commit para assegurar a durabilidade da operação
    err = tx.Commit(ctx)
    if err != nil {
        return fmt.Errorf("failed to commit transaction: %v", err)
    }

    fmt.Println("Transaction committed successfully.")
    return nil
}
```