# Go Balancer

## Instalación

Para compilar el proyecto se necesita el compilador de [go](https://golang.org/doc/install).

## Pasos para correr el load balancer

1. Ir al `src` del proyecto
2. Correr `go build -o go-balancer`
3. Correr `./go-balancer`

Nota: el `config.json` a usar debe estar en `/src`

## Pasos para correr el servidor test

1. Ir al `test` del proyecto
2. Correr `go build -o test`
3. Correr `./test 8081`

## Benchmark

Los benchmarks se encuentra en la carpeta `benchmark` y los resultados en el archivo `resultados.md`