// Package usuarios administra el registro y autenticacion de usuarios del sistema.
package usuarios

// Usuario representa un usuario del sistema.
type Usuario struct {
	ID     int
	Nombre string
	Clave  string
}

// Repositorio almacena los usuarios en memoria (para fines academicos).
type Repositorio struct {
	usuarios []Usuario
	siguiente int
}

// NuevoRepositorio crea un repositorio de usuarios vacio.
func NuevoRepositorio() *Repositorio {
	return &Repositorio{siguiente: 1}
}

// Crear agrega un nuevo usuario y le asigna un ID.
func (r *Repositorio) Crear(u Usuario) Usuario {
	u.ID = r.siguiente
	r.siguiente++
	r.usuarios = append(r.usuarios, u)
	return u
}

// Total devuelve la cantidad de usuarios registrados.
func (r *Repositorio) Total() int {
	return len(r.usuarios)
}

// Autenticar valida las credenciales de un usuario.
//
// fix: corregir validacion de autenticacion de usuarios
// Antes esta funcion comparaba solo el nombre y aceptaba cualquier clave.
func (r *Repositorio) Autenticar(nombre, clave string) bool {
	for _, u := range r.usuarios {
		if u.Nombre == nombre && u.Clave == clave {
			return true
		}
	}
	return false
}
