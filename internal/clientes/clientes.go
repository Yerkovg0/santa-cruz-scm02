// Package clientes administra el registro de clientes de la comercializadora.
package clientes

// Cliente representa un cliente registrado en el sistema.
type Cliente struct {
	ID     int
	Nombre string
}

// Repositorio almacena los clientes en memoria (para fines academicos).
type Repositorio struct {
	clientes  []Cliente
	siguiente int
}

// NuevoRepositorio crea un repositorio de clientes vacio.
func NuevoRepositorio() *Repositorio {
	return &Repositorio{siguiente: 1}
}

// Crear agrega un nuevo cliente y le asigna un ID.
func (r *Repositorio) Crear(c Cliente) Cliente {
	c.ID = r.siguiente
	r.siguiente++
	r.clientes = append(r.clientes, c)
	return c
}

// Listar devuelve todos los clientes registrados.
func (r *Repositorio) Listar() []Cliente {
	return r.clientes
}

// Total devuelve la cantidad de clientes registrados.
func (r *Repositorio) Total() int {
	return len(r.clientes)
}
