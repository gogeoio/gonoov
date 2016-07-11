# Changelog

## 0.X.X - Next

### Upcoming

## 0.1.3 - 2016-07-11

### Adicionado

- Função para criar tipo NoovTime
- Função string para tipo NoovString
- Tratamento da data de emissão entre (dEmi e dhEmi)

### Modificado

- Parâmetro de paginação para endpoint ```/app/nfe```

## 0.1.2 - 2016-07-06

### Adicionado

- Autenticar novamente apenas se tiver passado 5 minutos após o primeiro login
- Tipo para detalhe da NFe com tratamento de array e objeto

## 0.1.1 - 2016-07-05

### Adicionado

- Paginação para endpoint de NFe
- Tipo específico para forçar parser de número para string (NoovString)
- Tipo específico de time (NoovTime)
- Exemplo de utilização no README
- Arquivo VERSION

### Modificado

- Retorno do endpoint de NFe com informação de paginação e erros

### Corrigido

- Mapeamento de tipos de diversos atributos

## 0.1.0 - 2016-06-30

### Adicionado

- Client com login
- Endpoint para NFe