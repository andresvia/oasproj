package project

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoadProject(t *testing.T) {

	Convey("Dado un archivo descriptor de proyecto inválido", t, func() {
		Convey("Debe entrar en panico", func() {
			So(func() { LoadProject("../test-data/project/proyecto_invalido") }, ShouldPanic)
		})

	})

	Convey("Dado un archivo descriptor de proyecto inexistente", t, func() {
		Convey("Debe entrar en panico", func() {
			So(func() { LoadProject("../test-data/project/proyecto_inexsistente") }, ShouldPanic)
		})

	})

	Convey("Dado un projecto existente", t, func() {
		p := LoadProject("../test-data/project/proyecto_valido")
		Convey("Se reconocen sus campos", func() {
			So(p.Project_name, ShouldEqual, "Proyecto Válido")
			So(p.Project_description, ShouldEqual, "Este es un proyecto de prueba válido")
			So(p.Project_purpose, ShouldEqual, "Probar que al cargar este proyecto no se entra en pánico")
			So(p.Programming_language, ShouldEqual, "Go")
			So(p.Organizational_unit, ShouldEqual, "Interno")
			So(p.Package_dependencies, ShouldHaveSameTypeAs, []string{})
		})
		Convey("Tiene una representación textual correcta", func() {
			So(fmt.Sprintf("%s", p), ShouldNotBeEmpty)
			So(fmt.Sprintf("%s", p), ShouldNotStartWith, "TEMPLATE_ERROR=")
		})

	})

}

func TestMetadataExists(t *testing.T) {
	Convey("Dado un projecto", t, func() {
		Convey("Detectar existencia de metadatos", func() {
			So(MetadataExists("../test-data/project/proyecto_valido"), ShouldEqual, true)
		})
		Convey("Detectar usencia de metadatos", func() {
			So(MetadataExists("../test-data/project/proyecto_inexistente"), ShouldEqual, false)
		})
	})
}
