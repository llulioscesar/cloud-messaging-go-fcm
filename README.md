# FCM

1. Agrega el SDK

```shell
go get firebase.google.com/go/v4
```

2. Configura la variable de entorno `GOOGLE_APPLICATION_CREDENTIALS` con la ruta del archivo JSON que contiene la clave de
   tu cuenta de servicio


3. Inicializa el SDK, como se muestra a continuación:
```go
app, err := firebase.NewApp(context.Background(), nil)
if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
}
```