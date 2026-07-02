package productos

import "testing"

func TestCrearProducto(t *testing.T) {
	repo := NuevoRepositorio()
	p := repo.Crear(Producto{Nombre: "Arroz 1kg", Precio: 8.5})

	if p.ID != 1 {
		t.Errorf("se esperaba ID 1, se obtuvo %d", p.ID)
	}
	if repo.Total() != 1 {
		t.Errorf("se esperaba 1 producto, se obtuvo %d", repo.Total())
	}
}
