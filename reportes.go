// Package reportes genera indicadores para el dashboard administrativo.
package reportes

import (
	"fmt"

	"github.com/TU-USUARIO/santa-cruz-scm/internal/clientes"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/inventario"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/productos"
	"github.com/TU-USUARIO/santa-cruz-scm/internal/usuarios"
)

// ResumenGeneral construye un resumen con los indicadores principales del sistema.
// Corresponde a la funcionalidad "dashboard administrativo con reportes" (v1.2.0).
func ResumenGeneral(u *usuarios.Repositorio, c *clientes.Repositorio, p *productos.Repositorio, i *inventario.Repositorio) string {
	return fmt.Sprintf(
		"Resumen general:\n- Usuarios registrados: %d\n- Clientes registrados: %d\n- Productos registrados: %d",
		u.Total(), c.Total(), p.Total(),
	)
}
