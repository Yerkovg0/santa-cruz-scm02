package clientes

import "testing"

func TestCrearCliente(t *testing.T) {
	repo := NuevoRepositorio()
	c := repo.Crear(Cliente{Nombre: "Distribuidora Bolivia"})

	if c.ID != 1 {
		t.Errorf("se esperaba ID 1, se obtuvo %d", c.ID)
	}
	if repo.Total() != 1 {
		t.Errorf("se esperaba 1 cliente, se obtuvo %d", repo.Total())
	}
}
