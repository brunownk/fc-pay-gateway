FROM golang:1.24.2-alpine

# Instala dependências básicas (útil para desenvolvimento)
RUN apk add --no-cache git

# Cria usuário e diretório
RUN adduser -D -g '' appuser && \
    mkdir -p /app && \
    chown -R appuser:appuser /app

WORKDIR /app

# Copia o código Go (mantendo permissões)
COPY --chown=appuser:appuser . .

USER appuser
CMD ["sh", "-c", "while sleep 3600; do :; done"]
