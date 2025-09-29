package product

func (h *Handler) Count() (int64,
	error) {
	return h.service.Count()

}
