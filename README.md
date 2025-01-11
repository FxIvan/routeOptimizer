# OptiRoute

## Descripción
**OptiRoute** es un proyecto que resuelve el problema de optimización de rutas para una empresa de logística. Dado un conjunto de ubicaciones con pesos que representan la importancia de los paquetes, el objetivo es seleccionar las ubicaciones que maximicen la importancia total, respetando un límite en el número de entregas permitidas por día.

---

## Problema

Una empresa de logística tiene que entregar paquetes en diferentes ubicaciones de una ciudad. Cada ubicación tiene un identificador único y un peso asociado que representa la importancia del paquete a entregar. Sin embargo, el camión de reparto solo puede realizar una cierta cantidad de entregas por día debido a restricciones de tiempo.

Tu tarea es encontrar la mejor combinación de ubicaciones que maximice la importancia de los paquetes entregados, sin exceder el número máximo de entregas permitido en un solo día.

---

## Entrada

1. Una lista de ubicaciones, donde cada ubicación tiene:
   - `id`: Identificador único (entero o cadena).
   - `weight`: Importancia del paquete (entero positivo).
2. Un número entero `maxDeliveries` que representa el máximo de entregas que se pueden realizar por día.

---

## Salida

Una lista con los `id` de las ubicaciones seleccionadas que maximicen la suma de los pesos, respetando el límite de entregas.

---

## Restricciones

1. Puedes asumir que el número de ubicaciones no excede los 10,000.
2. Cada peso será un número entero positivo.
3. Debes resolver el problema de manera eficiente.

---

## Ejemplo

**Entrada:**
```json
locations = [
  { "id": "A", "weight": 10 },
  { "id": "B", "weight": 20 },
  { "id": "C", "weight": 15 },
  { "id": "D", "weight": 30 },
  { "id": "E", "weight": 25 }
]
maxDeliveries = 3
```

**Salida esperada:**
```json
["B", "D", "E"]
```

**Explicación:**
La combinación de ubicaciones que maximiza el peso total es `B`, `D`, y `E`, con un peso total de `20 + 30 + 25 = 75`.

---

## Tecnologías recomendadas
Puedes implementar este proyecto en:

- **Node.js**: Para un enfoque en JavaScript/TypeScript y su ecosistema.
- **Golang**: Para un enfoque en eficiencia y concurrencia.

---

## Siguientes pasos

1. Diseña el algoritmo para resolver el problema (puedes usar Programación Dinámica, Greedy, etc.).
2. Implementa el backend en el lenguaje de tu preferencia.
3. Escribe pruebas unitarias para garantizar la corrección de tu solución.
4. Documenta cómo ejecutar y probar tu proyecto.
