# go-sandbox-clean

```mermaid
---
title: Payment
---
classDiagram
    class Payment {
        +paymentID : int
    }
    class PaymentRepository{
        +GetPaymentID(limit int) (*[]Payment, error)
    }
    class paymentRepository {
        -db : *sql.DB
    }
    class PaymentController {
        -paymentRepository : PaymentRepository
    }
    Payment -- PaymentRepository : uses
    PaymentRepository <|.. paymentRepository : implements
    PaymentController -- PaymentRepository : uses
```
