# arca-grid
QUID Listen/Notify

## Que hace?
Esta parrilla de computo permite partir en pedazos un computo completo en _handlers_. Cada servidor es independiente y se supone que cada servidor debe conocer la estructura de los otros servidores a los cuales se pretende acceder.

Actualmente se utiliza un sistema simple basado en el patron _LISTEN/NOTIFY_ para intercambiar mensajes.

## Por hacer
- _Listen_ segun metodo. Ej. _Listen_ sobre _Query_, sobre _Update_ ...
- _Before_ y _After_
- Esquema sencillo de administración de transacciones
