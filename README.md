# Comandos de servidor

Estos comandos pueden sera accedidos una vez se intvita al bot a un servidor y comenzando a escrbir `/` en un canal al que el bot tenga acceso.

## Shuffle

### lista

Devuelve una lista de opciones en orden aleatoreo.

```
/shuffle lista [opciones]
```

- opciones: Lista de opciones separadas por espacios.

### rol

Devuelve una lista de los usuarios que tienen el rol indicado, en un orden aleatoreo.

```
/shuffle rol [rol]
```

- rol: Referencia a un rol del servidor.

### audio

Devuelve una lista de los usuarios conectados al canal indicado, en un orden aleatoreo.

```
/shuffle audio [canal]
```

- canal: Referencia a un canal de audio del servidor.


## Pick

### lista

Devuelve una cantidad de opciones aleatoreas de entre las seleccionadas.

```
/pick lista [opciones] [cantidad]
```

- opciones: Lista de opciones separadas por espacios.
- cantidad: Cantidad de opciones a elegir, por defecto es 1.

### rol

Devuelve una cantidad de usuarios aleatoreos que tengan el rol especificado

```
/pick rol [rol] [cantidad]
```

- rol: Referencia a un rol del servidor.
- cantidad: Cantidad de usuarios a devolver, por defecto es 1.

### audio

Devuelve una cantidad de usuarios aleatoreos que estén conectados al canal especificado.

```
/pick canal [canal] [cantidad]
```

- canal: Referencia a un canal de audio del servidor.
- cantidad: Cantidad de usuarios a devolver, por defecto es 1.

