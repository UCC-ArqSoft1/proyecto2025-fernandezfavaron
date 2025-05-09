package domain

import "time"

type Usuario struct {
	ID           uint   `gorm:"primaryKey" json:"id_usuario"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Nombre       string `json:"nombre"`
	Admin        bool   `gorm:"type:enum('socio','administrador')" json:"admin"`
	Foto         string `json:"foto"`
}

type Actividad struct {
	ID          uint      `gorm:"primaryKey" json:"id_actividad"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Dia         string    `json:"dia"`
	Horario     time.Time `json:"horario"`
	Duracion    int       `json:"duracion"`
	Cupos       int       `json:"cupos"`
	Categoria   string    `json:"categoria"`
	Instructor  string    `json:"instructor"`
	FotoURL     string    `json:"fotourl"`
}

type Inscripcion struct {
	ID               uint   `gorm:"primaryKey" json:"id_inscripcion"`
	IDUsuario        uint   `json:"id_usuario"`
	IDActividad      uint   `json:"id_actividad"`
	FechaInscripcion string `gorm:"autoCreateTime" json:"fecha_inscripcion"`
}
