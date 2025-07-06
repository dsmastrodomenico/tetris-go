# Usa una imagen base de Go que incluya las herramientas de compilación
FROM golang:1.23.10-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum para descargar las dependencias primero.
# Esto aprovecha el cache de Docker si las dependencias no cambian.
COPY go.mod ./
COPY go.sum ./

# Descarga las dependencias del módulo
RUN go mod download

# Copia todo el código fuente de tu proyecto al contenedor
COPY . .

# Compila la aplicación. -o tetris-go especifica el nombre del ejecutable.
RUN go build -o tetris-go .

# Comando por defecto para ejecutar la aplicación cuando el contenedor se inicie
CMD ["./tetris-go"]