# arca-grid
QUID Listen/Notify

## Que hace?
Esta parrilla de computo permite partir en pedazos un computo completo en _handlers_. Cada servidor es independiente y se supone que cada servidor debe conocer la estructura de los otros servidores a los cuales pretende acceder.

Actualmente se utiliza un sistema simple basado en el patrón _LISTEN/NOTIFY_ para intercambiar mensajes.

## Ejemplos
Ver la carpeta demo

## Por hacer
- _Listen_ según método. Ej. _Listen_ sobre _Query_, sobre _Update_ ...
- _Unlisten_ general
- _Unlisten_ según método
- Pruebas sobre _notify_ cuando un listener ha cesado
- Tornar ésta solución en un módulo