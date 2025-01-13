# clan-de-raid

## Descripción
Descripción del proyecto aquí.

## Estructura
```
project-name/
├── cmd/
│   └── api/
│       └── main.go          # Punto de entrada de la aplicación
├── config/
│   └── config.go           # Configuraciones de la aplicación
├── internal/
│   ├── app/
│   │   └── server.go       # Configuración del servidor HTTP
│   ├── db/                 # conexion a la base de datos
│   ├── handlers/           # Manejadores de las rutas HTTP
│   ├── middleware/         # Middleware personalizados
│   ├── models/             # Definición de estructuras de datos
│   ├── policies/           # Control de acceso o permisos
│   ├── repositories/       # Capa de acceso a datos
│   ├── routes/            # Capa de acceso a datos
│   ├── services/          # Lógica de negocio
│   ├── requests/          # Estructuras para validación de requests
│   └── resources/         # Transformadores de respuesta (DTOs)
├── pkg/                    # Código que puede ser reutilizado por otros proyectos
│   └── validator/
├── storage/               # Archivos generados (logs, caché, etc.)
├── tests/                 # Pruebas de integración
├── .env                   # Variables de entorno
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

