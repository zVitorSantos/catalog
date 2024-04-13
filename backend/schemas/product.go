package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Ref          string
	Imagem       string
	Descricao    string
	Categoria_1  string
	Categoria_2  string
	Categoria_3  string
	Complementos string
	Material     string
	Peso         int64
	Altura       int64
	Largura      int64
	Comprimento  int64
	Valor        int64
	Matriz       bool
	Piloto       bool
	Desenho      bool
}

type ProductResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	Ref          string    `json:"ref"`
	Imagem       string    `json:"imagem"`
	Descricao    string    `json:"descricao"`
	Categoria_1  string    `json:"categoria_1"`
	Categoria_2  string    `json:"categoria_2"`
	Categoria_3  string    `json:"categoria_3"`
	Complementos string    `json:"complementos"`
	Material     string    `json:"material"`
	Peso         int64     `json:"peso"`
	Altura       int64     `json:"altura"`
	Largura      int64     `json:"largura"`
	Comprimento  int64     `json:"comprimento"`
	Valor        int64     `json:"valor"`
	Matriz       bool      `json:"matriz"`
	Piloto       bool      `json:"piloto"`
	Desenho      bool      `json:"desenho"`
}
