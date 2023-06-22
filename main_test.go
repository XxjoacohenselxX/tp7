package main

import "testing"

func TestValidar(t *testing.T) {

	tarjetaPruebaMastercard := 5031755734530604
	tarjetaPruebaVisa := 4509953566233704
	tarjetaPruebaAmex := 371180303257522

	tarjetaPruebaMastercardInvalida := 5031755734530601
	tarjetaPruebaVisaInvalida := 4509953536233704
	tarjetaPruebaAmexInvalida := 771180303257522

	result := validar(tarjetaPruebaMastercard)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaVisa)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaAmex)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaMastercardInvalida)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, false)
	}

	result = validar(tarjetaPruebaVisaInvalida)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, false)
	}

	result = validar(tarjetaPruebaAmexInvalida)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, false)
	}

}
