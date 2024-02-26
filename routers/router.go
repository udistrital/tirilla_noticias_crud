// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/tirilla_noticias_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/noticia",
			beego.NSInclude(
				&controllers.NoticiaController{},
			),
		),

		beego.NSNamespace("/noticia_tipo_etiqueta",
			beego.NSInclude(
				&controllers.NoticiaTipoEtiquetaController{},
			),
		),

		beego.NSNamespace("/tipo_etiqueta",
			beego.NSInclude(
				&controllers.TipoEtiquetaController{},
			),
		),

		beego.NSNamespace("/modulo_publicacion",
			beego.NSInclude(
				&controllers.ModuloPublicacionController{},
			),
		),

		beego.NSNamespace("/noticia_tipo_contenido",
			beego.NSInclude(
				&controllers.NoticiaTipoContenidoController{},
			),
		),

		beego.NSNamespace("/tipo_contenido",
			beego.NSInclude(
				&controllers.TipoContenidoController{},
			),
		),

		beego.NSNamespace("/tipo_estilo",
			beego.NSInclude(
				&controllers.TipoEstiloController{},
			),
		),

		beego.NSNamespace("/tipo_prioridad",
			beego.NSInclude(
				&controllers.TipoPrioridadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
