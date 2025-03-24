# Usar una imagen de Go más reciente (1.20 o superior)() se cambia a 1.21
FROM golang:1.23 AS builder


# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

# Copiar archivos de módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# **Construir el binario durante la construcción de la imagen**
RUN go build -o app

# Asegurar que el binario tenga permisos de ejecución


# Exponer el puerto que utiliza la aplicación
EXPOSE 8000

# Comando para ejecutar CompileDaemon y recompilar automáticamente en cambios
CMD ["sh", "-c", "CompileDaemon --build='go build -o app .' --command=./app"]

