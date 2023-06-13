package models

type Mujer struct {
	Hombre
}

func (mujer *Mujer) Respirar() {
	mujer.respirando = true
}
func (mujer *Mujer) Comer() {
	mujer.comiendo = true
}
func (mujer *Mujer) Pensar() {
	mujer.pensando = true
}
func (mujer *Mujer) Sexo() string {
	return "Mujer"

}
