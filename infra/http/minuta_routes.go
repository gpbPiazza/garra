package http

func (r *router) SetMinutaRoutes() {
	r.apiV1.Post("/generator/minuta", PostGeneratorMinutaHandler)
}
