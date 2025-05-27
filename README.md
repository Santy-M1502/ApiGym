# API Lista de Ejercicios

API REST para gestionar una lista personal de ejercicios. Permite a los usuarios registrarse, iniciar sesión, agregar y eliminar ejercicios.

---

## Tecnologías

- Node.js
- Express
- JWT para autenticación
- Base de datos (configurable)

---

## Funcionalidades

- Registro de usuario
- Login y autenticación con JWT
- Agregar ejercicios a la lista personal
- Eliminar ejercicios de la lista
- Listar ejercicios del usuario autenticado

---

## Endpoints

### Autenticación

| Método | Ruta       | Descripción            | Cuerpo (JSON)                         |
| ------ | ---------- | --------------------- | ----------------------------------- |
| POST   | `/register`| Registrar usuario      | `{ "username": "user", "password": "pass" }` |
| POST   | `/login`   | Login de usuario       | `{ "username": "user", "password": "pass" }` |

### Ejercicios (requieren token JWT en header Authorization)

| Método | Ruta            | Descripción                | Cuerpo (JSON)                 |
| ------ | --------------- | --------------------------| -----------------------------|
| POST   | `/exercises`    | Agregar nuevo ejercicio    | `{ "name": "Flexiones", "reps": 15 }` |
| DELETE | `/exercises/:id`| Eliminar ejercicio por ID  | N/A                           |
| GET    | `/exercises`    | Listar ejercicios del usuario | N/A                         |

---

## Autenticación

Para endpoints protegidos enviar el token JWT en el header:
Authorization: Bearer <TOKEN>
