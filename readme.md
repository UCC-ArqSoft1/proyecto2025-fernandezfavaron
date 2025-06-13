# 🏗️ Sistema de Gestión de Actividades Deportivas

Este proyecto es una API REST para gestionar actividades deportivas, socios y administradores. Permite a los usuarios ver y registrarse en actividades, mientras que los administradores pueden crear, editar o eliminar actividades.

## 📚 Descripción

El sistema está diseñado para manejar diferentes tipos de usuarios (socios y administradores) y brindar funcionalidades de autenticación, visualización y gestión de actividades.
# Como correr el proyecto 
Usar uno de los siguientes comandos dependiendo el sistema operativo el que se encuentre para correr el proyecto, debe colocar el usuario, contraseña y el nombre de la base de datos
	Linux
	export DB_DSN="usuario:miclave123@tcp(127.0.0.1:3306)/basedatosnombre?charset=utf8mb4&parseTime=True&loc=Local"
	Windows Powershell
	setx DB_DSN "usuario:miclave123@tcp(127.0.0.1:3306)/basedatosnombre?charset=utf8mb4&parseTime=True&loc=Local"
 Luego abrir dos terminales, abrir una terminal para el backend y otra para el frontend en sus respectivas carpetas,
 En el backend ejecutar: go run main.go
 En el frontend ejecutar: 
 npm install
 npm run dev

## 🚀 Endpoints

### 🔐 Autenticación

- `POST /login`:  
  Endpoint de autenticación. Recibe credenciales (`email` y `password`) y devuelve un token JWT.

- `POST /register`:  
  Registra un nuevo usuario. Recibe (`email`, `password`, `nombre`) y devuelve un token JWT.

### 🏃‍♂️ Actividades

- `GET /actividades`:  
  Devuelve un listado de actividades disponibles. Se pueden usar parámetros para buscar.

- `GET /actividades/:id`:  
  Devuelve los detalles de una actividad específica por su ID.

- `GET /misactividades`:  
  Devuelve las actividades en las que el usuario autenticado está inscrito. Requiere token JWT.

### 📝 Inscripciones

- `POST /inscripcion`:  
  Inscribe al usuario autenticado en una actividad. Recibe en el body: `id_actividad`. Requiere token JWT.

- `POST /unscripcion`:  
  Elimina la inscripción del usuario autenticado en una actividad. Recibe en el body: `id_actividad`. Requiere token JWT.
  
### 🛠️ Administración

- `PUT /admin/actividades/{id}`:  
  Permite a un administrador editar una actividad existente. Requiere autenticación y permisos de administrador.

## 🧱 Estructura de Datos

### Usuario

```go
type Usuario struct {
    ID           uint      `gorm:"primaryKey" json:"id_usuario"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password"`
    Nombre       string    `json:"nombre"`
    IsAdmin      bool      `gorm:"type:bool" json:"admin"`
    Foto         string    `json:"foto"`
    Usuarios     []Usuario `gorm:"many2many:inscripcions;" json:"usuarios"`
}
```

### Actividad

```go
type Actividad struct {
    ID          uint        `gorm:"primaryKey" json:"id_actividad"`
    Nombre      string      `json:"nombre"`
    Descripcion string      `json:"descripcion"`
    Dia         string      `json:"dia"`
    Horario     time.Time   `json:"horario"`
    Duracion    int         `json:"duracion"`
    Cupos       int         `json:"cupos"`
    Categoria   string      `json:"categoria"`
    Instructor  string      `json:"instructor"`
    FotoURL     string      `json:"fotourl"`
    Actividades []Actividad `gorm:"many2many:inscripcions;" json:"actividades"`
}
```

### Inscripcion

```go
type Inscripcion struct {
    IDUsuario        uint      `gorm:"primaryKey" json:"id_usuario"`
    IDActividad      uint      `gorm:"primaryKey" json:"id_actividad"`
    Usuario          Usuario   `gorm:"foreignKey:IDUsuario;constraint:OnDelete:CASCADE" json:"-"`
    Actividad        Actividad `gorm:"foreignKey:IDActividad;constraint:OnDelete:CASCADE" json:"-"`
    FechaInscripcion time.Time `gorm:"autoCreateTime" json:"fecha_inscripcion"`
}
```
# Como correr el proyecto 
