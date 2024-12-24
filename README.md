# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación
Ejecuta el siguieente comando para instalar el paquete:
```bash
go get -u github.com/lautaromdelgado/greetings
```

## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
	"fmt"
	"log"

	"github.com/lautaromdelgado/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if msg, err := greetings.Hello("Javier Mieli"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}

	var names = []string{"Ana", "Milagros", "Agustina", "Valentina"}

	if msg, err := greetings.Hellos(names); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}
}
 
```