// Package productos administra el catalogo de productos de la comercializadora.
package productos

// Producto representa un producto del catalogo.
type Producto struct {
	ID     int
	Nombre string
	Precio float64
}

// Repositorio almacena los productos en memoria (para fines academicos).
type Repositorio struct {
	productos []Producto
	siguiente int
}

// NuevoRepositorio crea un repositorio de productos vacio.
func NuevoRepositorio() *Repositorio {
	return &Repositorio{siguiente: 1}
}

// Crear agrega un nuevo producto y le asigna un ID.
func (r *Repositorio) Crear(p Producto) Producto {
	p.ID = r.siguiente
	r.siguiente++
	r.productos = append(r.productos, p)
	return p
}

// Listar devuelve todos los productos registrados.
func (r *Repositorio) Listar() []Producto {
	return r.productos
}

// Total devuelve la cantidad de productos registrados.
func (r *Repositorio) Total() int {
	return len(r.productos)
}
