1. Escribir una función `esMasLargo` que reciba dos strings y retorne `true` si el primer string es más largo que el segundo, y `false` en caso contrario.

2. Escribir una función `unirConY` que reciba dos strings y retorne un string que resulte de unir ambos con un " y " en el medio. 
Por ejemplo:
    `unirConY("pan","queso")` da como resultado `"pan y queso"`.

3. Escribir una función `esPalindromo` que recibe un string y retorna `true` si es palíndromo, `false` en caso contrario.

4. Escribir una función `contarLetra` que reciba un caracter y un string y retorne el número de veces que ese caracter aparece en el string.

5. Escribir una función `contarPalabra` que reciba un string palabra y un string frase, y retorne el número de veces que esa palabra aparece en la frase.

6. Escribir una función `reemplazarPalabra` que reciba una frase, una palabra a ser reemplazada y una palabra con la cual reemplazarla, y retorne el string que resulte de reemplazar una palabra por otra en la frase dada. 
Por ejemplo: `reemplazarPalabra("Usted es Homero Simpson", "Simpson", "Thompson")` retorna `"Usted es Homero Thompson"`

7. Repasar del libro cómo trabajar con funciones variádicas. Escribir una función `enumerar` que reciba una cantidad variable de strings y retorne el resultado de unirlos con comas e `y`. 
Por ejemplo: 
    `enumerar("uno","dos","tres")` debe retornar `"uno, dos y tres"`.
    `enumerar("ancho","alto")` debe retornar `"ancho  alto"`.
    `enumerar("rojo","azul","amarillo", "negro")` debe retornar `"rojo, azul, amarillo y negro"`.

7.5 Dada una lista de números, encontrar el mínimo.

8. Escribir una función `ordenar` que reciba un slice de números y retorne un slice con esos números en orden creciente.

9. Escribir una función `combinar` que reciba dos arrays de números ya ordenados de manera creciente y retorne un slice que contenga los números de ambos slices, ordenados de manera creciente.
Nota: cuando dependemos de que un argumento ya cumpla determinadas reglas, decimos que se trata de una "precondición". En este caso nuestra'precondición es que nuestros argumentos ya vengan ordenados.