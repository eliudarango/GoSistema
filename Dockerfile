# Usa la imagen oficial de Go como base
FROM golang:1.24.1-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de la aplicación Go al contenedor
COPY . .

# Instala las dependencias de Go (si las tienes)
RUN go mod tidy

# Exponer el puerto 8080 (asegúrate de que tu aplicación esté escuchando en este puerto)
EXPOSE 8080

# Comando para ejecutar la aplicación Go
CMD ["go", "run", "main.go"]
