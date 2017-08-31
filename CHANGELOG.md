# Changelog

## 0.X.X - Next

### Upcoming

### 0.3.0

### Adicionado

- Serviço de totalizadores

### 0.2.0 - 2017-07-14

### Adicionado

- Serviço de stats de notas fiscais emitidas

### 0.1.6 - 2017-04-28

### Adicionado

- Adicionar campo CPF para NfeDest
- Melhorias no parse de date/time

## 0.1.5 - 2016-10-27

### Adicionado

- Retry quando ocorrer erro 503
- Unmarshall customizado para campo ```NFe.infNFe.infAdic.obsCont```

## 0.1.4 - 2016-10-20

### Adicionado

- Unmarshall customizado para campo ```NFe.infNFe.transp.vol``` para tratar casos que o dado bruto está como array
- Dado raw retornado pelo servidor para facilitar validação/debug
- Timeout para http.Client

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