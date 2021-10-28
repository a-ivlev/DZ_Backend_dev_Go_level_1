FROM golang:1.17.2 as builder
# Создаём директорию build и переходим в неё.
WORKDIR /app
# Копируем все файлы из директории ./cmd локальной машины в директорию /rental образа.
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Запускаем компиляцию программы на go и сохраняем полученный бинарный файл server в директорию /rental/ образа.
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../rental/server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./server ./rental/cmd/stress-test



FROM scratch
# Копируем бинарник server из образа builder в корневую директорию.
COPY --from=builder /rental/server /
# Информационная команда показывающая на каком порту будет работать приложение.
EXPOSE 8080
# Запускаем приложение.
ENTRYPOINT ["/server"]
