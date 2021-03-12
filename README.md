# Сервис, возвращающий срез последовательности чисел из ряда Фибоначчи.

Не забудьте добавить файл .env с примера env.local или env.docker

### Для запуска приложения с  docker:

```
make build && make run
```

### Для запуска тестов:

```
make test
```

### Для получения чисел Фибоначии по протоколу http, сделайте GET запрос:

```
localhost:8080/api/get-fibonacci?x=0&y=90
```

### Для получения чисел Фибоначии по протоколу grpc, необходимо установить клиент [evans](https://github.com/ktr0731/evans). И указать proto файл .proto/fibonacci.proto

```
evans ./proto/fibonacci.proto -p 4040

call FibonacciSlice
x (TYPE_UINT64) => 2
y (TYPE_UINT64) => 4
```
