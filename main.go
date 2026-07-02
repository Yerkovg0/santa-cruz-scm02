// Comando principal del sistema de la Comercializadora Santa Cruz S.R.L.
package main

import (
	"fmt"

	"github.com/TU-USUARIO/santa-cruz-scm/internal/clientes"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/inventario"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/productos"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/reportes"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/usuarios"
)

func main() {
	fmt.Println("Sistema de Informacion - Comercializadora Santa Cruz S.R.L.")

	// Ejemplo minimo de uso de cada modulo (evidencia funcional del sistema)
	usuariosRepo := usuarios.NuevoRepositorio()
	usuariosRepo.Crear(usuarios.Usuario{Nombre: "admin", Clave: "1234"})

	clientesRepo := clientes.NuevoRepositorio()
	clientesRepo.Crear(clientes.Cliente{Nombre: "Distribuidora Bolivia"})

	productosRepo := productos.NuevoRepositorio()
	p := productosRepo.Crear(productos.Producto{Nombre: "Arroz 1kg", Precio: 8.5})

	inventarioRepo := inventario.NuevoRepositorio()
	inventarioRepo.Ingresar(p.ID, 100)

	fmt.Println(reportes.ResumenGeneral(usuariosRepo, clientesRepo, productosRepo, inventarioRepo))
}
