---
sidebar_position: 1
title: Autenticação
---
# Autenticação na aplicação

A autenticação na aplicação é feita através de um sistema de autenticação baseado em tokens. O token é gerado pelo backend e enviado para a aplicação, que o guarda no FlutterSecureStorage ou no LocalStorage no caso da aplicação web. O token é utilizado para autenticar o utilizador em todas as chamadas ao backend.

```mermaid
sequenceDiagram
    actor User
    participant App
    participant Backend

    User->>App: Login with email and password
    App->>Backend: Send gRPC request
    Backend->>Backend: Random wait (+- 1s)
    Backend->>App: Send authentication token
    App-->>App: Save Auth Token to FlutterSecureStorage
    App->>User: Success
```

# Registo na aplicação

O utilizador deve registar-se com o nome, email e password. Poderá também adicionar

```mermaid
sequenceDiagram
    participant User
    participant App
    participant Backend
    participant Backblaze

    User->>App: Fills Form (Name, Email, Password, Photo, Interests and ContractTypes)
    App->>Backend: Send Registration Data (Name, Email, Password, Photo, Interests and ContractTypes)
    Backend->>Backend: Validate Registration Data
    Backend->>Backend: Store User Details (Name, Email, Password, Interests and ContractTypes)
    opt has photo
        Backend->>Backblaze: Request Presigned Upload URL
        Backblaze-->>Backend: Return Presigned Upload URL
        Backend->>App: Send Presigned Upload URL
        App->>Backblaze: Upload Photo using Presigned URL
        Note over Backblaze: Photo Uploaded, Event Triggered
        Backblaze-->>Backend: Event Notification (Photo Uploaded)
        Backend->>Backend: Save Photo object key to User
    end
    Backend->>App: Registration Successful
    App->>User: Registration Successful
```