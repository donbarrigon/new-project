# new-project

## Descripción
new-projectes es un framework para desarrollar APIs RESTful con Go.
esta arquitectura se basa en el patrón MVC (Modelo-Vista-Controlador). 
utiliza mariaDB como base de datos pero también 
puede utilizar PostgreSQL mas o menos ya que  aun no he agregado el soporte para postgre pero muchas cosas funcionan

aun esta en fase de desarrollo.
por favor pruebalo y dame cualquier feedback.

## Estructura
```
project-name/
├── cmd/
│   └── api/
│       └── main.go         # Punto de entrada de la aplicación
├── config/
│   └── config.go           # Configuraciones de la aplicación
├── internal/
│   ├── app/
│   │   └── server.go       # Configuración del servidor HTTP
│   ├── database/           # migraciones y seeders
│   ├── controller/         # Manejadores de las rutas HTTP
│   ├── middleware/         # Middleware personalizados
│   ├── models/             # Definición de estructuras de datos
│   ├── orm/                # conexion a la base de datos y consultas
│   ├── policies/           # Control de acceso o permisos
│   ├── repositories/       # Capa de acceso a datos
│   ├── routes/             # Enrutador
│   ├── services/           # Lógica de negocio
│   ├── requests/           # Estructuras para validación de requests
│   └── resources/          # Transformadores de respuesta (DTOs)
├── pkg/                    # Código que puede ser reutilizado por otros proyectos
│   └── formatter/          # funciones de formateo de datos
│   └── validator/          # funciones de validacion
├── storage/                # Archivos generados (logs, caché, etc.)
├── tests/                  # Pruebas de integración
├── .env                    # Variables de entorno
├── .gitignore
├── go.mod
└── README.md
```

## Instalación

```bash
go mod download
```

## Uso

```bash
go run cmd/api/main.go
```

