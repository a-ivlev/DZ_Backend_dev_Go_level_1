FROM golang:1.17.2 as builder
# Создаём директорию build и переходим в неё.
WORKDIR /app

# Копируем файлы go.mod и go.sum и делаем загрузку, чтобы вовремя пересборки контейнера зависимости
# подтягивались из слоёв.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Копируем все файлы из директории ./rental локальной машины в директорию /app/rental образа.
COPY ./rental ./rental

# Запускаем компиляцию программы на go и сохраняем полученный бинарный файл server в директорию /rental/ образа.
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../rental/server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../stress-test ./rental/cmd/stress-test



FROM scratch

WORKDIR /app

# Копируем бинарник server из образа builder в корневую директорию.
COPY --from=builder /stress-test /

# Копируем сертификаты и таймзоны.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# Устанавливаем в переменную окружения свою таймзону.
ENV TZ=Europe/Moscow

# Информационная команда показывающая на каком порту будет работать приложение.
EXPOSE 8080

# Устанавливаем по дефолту переменные окружения, которые можно переопределить при запуске контейнера.
ENV POSTGRES_DSN=postgres://manager:qwerty@localhost:5432/rental_db

# Запускаем приложение.
CMD ["/stress-test"]
