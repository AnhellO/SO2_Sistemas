# Práctica 2 - 2do Parcial

Manejo de errores e interrupciones en Go

## errores.go

Crear un programa utilizando Go, en donde lean dos números `N` y `M` por medio de la consola, y para esos dos números calcules una divisón `N / M`. El programa deberá de leer ambos números hasta que estos sean 0 (esto quiere decir; detener la ejecución del programa cuando esta condición se cumpla). Si `M` es igual a 0, entonces implementa el manejo de errores para que el programa imprima un mensaje de advertencia y continue la ejecución. En caso de que `M != 0`, entonces imprime el resultado de la división a la consola y continúa con el programa.

## Entrada

Dos números `N` y `M`

## Salida

El resultado de la división de `N / M`, o un mensaje `"No te pases, M es 0!"` en caso de que `M == 0`

## Ejemplo

### Entrada
```
3 3
18 6
4 2
4 0
0 0
```

### Salida
```
1
3
2
No te pases, M es 0!
```

## Referencias
- [Errores en Go](https://golang.org/pkg/errors/)
- [Mejores prácticas para manejo de errores en Go](https://medium.com/@sebdah/go-best-practices-error-handling-2d15e1f0c5ee)