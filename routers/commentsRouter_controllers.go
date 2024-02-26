package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:ModuloPublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoContenidoController"],
        beego.ControllerComments{
            Method: "GetContenidosByIdNoticia",
            Router: "/contenido/:idNoticia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:NoticiaTipoEtiquetaController"],
        beego.ControllerComments{
            Method: "GetEtiquetasByIdNoticia",
            Router: "/etiquetas/:idNoticia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoContenidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEstiloController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoEtiquetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/tirilla_noticias_crud/controllers:TipoPrioridadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
