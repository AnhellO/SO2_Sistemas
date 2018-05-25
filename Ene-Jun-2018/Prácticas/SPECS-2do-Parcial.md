# 2do. Parcial

100% Práctico! `:D`

## 2do.go

Crear un programa utilizando Go, en donde simulen un restaurante. Deberán de leer la entrada de datos a través de la consola. El programa terminará hasta que se le indique con una opción en específico. Las siguientes opciones son las que deberán de manejar en el programa:

1. Pedir comida
  - Al elegir la opción `1`, deberán de leer un string que represente a la comida que quieren pedir
  - Ejemplo: `pizza`
  - Si hay pizza suficiente, entonces lo indican en la consola con algo como `Hay pizza! :D`, decrementan la cantida de pizzas, y continuan con el programa; en caso de que no haya pizza, imprimen algo como `No hay pizza! :C` y continuan con el programa
2. Surtir comida
  - Al elegir la opción `2`, deberán de leer un string que represente a la comida que quieren surtir, seguido por un espacio, y luego por un número, que representa la cantidad de piezas de comida que van a surtir
  - Ejemplo: `hamburguesa 2`
  - Después imprimen a consola un mensaje como el siguiente `Comida surtida con éxito!`, en caso de que se haya podido surtir la comida con éxito, o un mensaje tipo `Fallo al surtir la comida!` en caso contrario. Después continuan con el programa
3. Imprimir comida
  - Al elegir la opción `3`, imprimen a la consola toda la comida que hay disponible y continuan con el programa
4. Salir
  - Terminan por completo con la ejecución del programa

## Puntos extra
- Código legible e indentado
- D.R.Y.