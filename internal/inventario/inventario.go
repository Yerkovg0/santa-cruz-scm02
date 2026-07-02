// Package inventario controla el stock disponible de cada producto.
package inventario

// Movimiento representa un ingreso o salida de stock.
type Movimiento struct {
	ProductoID int
	Cantidad   int // positivo = ingreso, negativo = salida
}

// Repositorio controla el stock de productos en memoria.
type Repositorio struct {
	movimientos []Movimiento
}

// NuevoRepositorio crea un repositorio de inventario vacio.
func NuevoRepositorio() *Repositorio {
	return &Repositorio{}
}

// Ingresar registra una entrada de stock para un producto.
func (r *Repositorio) Ingresar(productoID, cantidad int) {
	r.movimientos = append(r.movimientos, Movimiento{ProductoID: productoID, Cantidad: cantidad})
}

// Retirar registra una salida de stock si hay disponibilidad suficiente.
// Devuelve false si no hay stock suficiente.
func (r *Repositorio) Retirar(productoID, cantidad int) bool {
	if r.StockDisponible(productoID) < cantidad {
		return false
	}
	r.movimientos = append(r.movimientos, Movimiento{ProductoID: productoID, Cantidad: -cantidad})
	return true
}

// StockDisponible calcula el stock actual de un producto.
//
// fix: corregir calculo de stock disponible
// La version anterior solo sumaba los ingresos y no restaba las salidas,
// lo que generaba un stock disponible incorrecto (siempre mayor al real).
func (r *Repositorio) StockDisponible(productoID int) int {
	total := 0
	for _, m := range r.movimientos {
		if m.ProductoID == productoID {
			total += m.Cantidad
		}
	}
	return total
}
