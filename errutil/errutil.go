package errutil

// Retorna el primer error del arreglo de errores que se pasa o <nil> si no hay ningún error.
func FirstOrNil(errs []error) (err error) {
	if len(errs) == 0 {
		return nil
	} else {
		return errs[0]
	}
}
