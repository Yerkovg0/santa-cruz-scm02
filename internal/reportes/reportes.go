// Package reportes genera indicadores para el dashboard administrativo.
package reportes

import (
	"fmt"

	"github.com/Yerkovg0/santa-cruz-scm02/internal/clientes"
	"github.com/Yerkovg0/santa-cruz-scm02/internal/inventario"
	"github.com/Yerkovg0/santa-cruz-scm02/internal/productos"
	"github.com/Yerkovg0/santa-cruz-scm02/internal/usuarios"
)

// ResumenGeneral construye un resumen con los indicadores principales del sistema.
// Corresponde a la funcionalidad "dashboard administrativo con reportes" (v1.2.0).
func ResumenGeneral(u *usuarios.Repositorio, c *clientes.Repositorio, p *productos.Repositorio, i *inventario.Repositorio) string {
	return fmt.Sprintf(
		"Resumen general:\n- Usuarios registrados: %d\n- Clientes registrados: %d\n- Productos registrados: %d",
		u.Total(), c.Total(), p.Total(),
	)
}
