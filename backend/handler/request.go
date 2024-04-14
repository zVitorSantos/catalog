package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateProduct

type CreateProductRequest struct {
	Ref          string `json:"ref"`
	Imagem       string `json:"imagem"`
	Descricao    string `json:"descricao"`
	Categoria_1  string `json:"categoria_1"`
	Categoria_2  string `json:"categoria_2"`
	Categoria_3  string `json:"categoria_3"`
	Complementos string `json:"complementos"`
	Material     string `json:"material"`
	Peso         int64  `json:"peso"`
	Altura       int64  `json:"altura"`
	Largura      int64  `json:"largura"`
	Comprimento  int64  `json:"comprimento"`
	Valor        int64  `json:"valor"`
	Matriz       *bool  `json:"matriz"`
	Piloto       *bool  `json:"piloto"`
	Desenho      *bool  `json:"desenho"`
}

func (r *CreateProductRequest) Validate() error {
	if r.Ref == "" && r.Imagem == "" && r.Descricao == "" && r.Categoria_1 == "" && r.Categoria_2 == "" && r.Categoria_3 == "" && r.Complementos == "" && r.Material == "" && r.Peso <= 0 && r.Altura <= 0 && r.Largura <= 0 && r.Comprimento <= 0 && r.Valor < 0 && r.Matriz == nil && r.Piloto == nil && r.Desenho == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Ref == "" {
		return errParamIsRequired("ref", "string")
	}
	if r.Imagem == "" {
		return errParamIsRequired("imagem", "string")
	}
	if r.Descricao == "" {
		return errParamIsRequired("descricao", "string")
	}
	if r.Categoria_1 == "" {
		return errParamIsRequired("categoria_1", "string")
	}
	if r.Categoria_2 == "" {
		return errParamIsRequired("categoria_2", "string")
	}
	if r.Categoria_3 == "" {
		return errParamIsRequired("categoria_3", "string")
	}
	if r.Complementos == "" {
		return errParamIsRequired("complementos", "string")
	}
	if r.Material == "" {
		return errParamIsRequired("material", "string")
	}
	if r.Peso <= 0 {
		return errParamIsRequired("peso", "int64")
	}
	if r.Altura <= 0 {
		return errParamIsRequired("altura", "int64")
	}
	if r.Largura <= 0 {
		return errParamIsRequired("largura", "int64")
	}
	if r.Comprimento <= 0 {
		return errParamIsRequired("comprimento", "int64")
	}
	if r.Valor < 0 {
		return errParamIsRequired("valor", "int64")
	}
	if r.Matriz == nil {
		return errParamIsRequired("matriz", "bool")
	}
	if r.Piloto == nil {
		return errParamIsRequired("piloto", "bool")
	}
	if r.Desenho == nil {
		return errParamIsRequired("desenho", "bool")
	}
	return nil
}

type UpdateProductRequest struct {
	Ref          string `json:"ref"`
	Imagem       string `json:"imagem"`
	Descricao    string `json:"descricao"`
	Categoria_1  string `json:"categoria_1"`
	Categoria_2  string `json:"categoria_2"`
	Categoria_3  string `json:"categoria_3"`
	Complementos string `json:"complementos"`
	Material     string `json:"material"`
	Peso         int64  `json:"peso"`
	Altura       int64  `json:"altura"`
	Largura      int64  `json:"largura"`
	Comprimento  int64  `json:"comprimento"`
	Valor        int64  `json:"valor"`
	Matriz       *bool  `json:"matriz"`
	Piloto       *bool  `json:"piloto"`
	Desenho      *bool  `json:"desenho"`
}

func (r *UpdateProductRequest) Validate() error {
	// If any field is provided, validation is thuthy
	if r.Ref == "" && r.Imagem == "" && r.Descricao == "" && r.Categoria_1 == "" && r.Categoria_2 == "" && r.Categoria_3 == "" && r.Complementos == "" && r.Material == "" && r.Peso <= 0 && r.Altura <= 0 && r.Largura <= 0 && r.Comprimento <= 0 && r.Valor < 0 && r.Matriz == nil && r.Piloto == nil && r.Desenho == nil {
		return nil
	}

	// If no field is provided, return error
	return fmt.Errorf("at least one valid field must be provided")
}
