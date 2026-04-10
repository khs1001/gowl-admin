package responses

type AdminRoute struct {
	Name      string          `json:"name"`
	Path      string          `json:"path"`
	Component string          `json:"component"`
	IsHome    int             `json:"is_home"`
	IframeUrl string          `json:"iframe_url"`
	UrlType   int             `json:"url_type"`
	KeepAlive int             `json:"keep_alive"`
	IsFull    int             `json:"is_full"`
	IsLink    bool            `json:"is_link"`
	PageSign  string          `json:"page_sign"`
	Meta      *AdminRouteMeta `json:"meta"`
	Children  []*AdminRoute   `json:"children"`
}

type AdminRouteMeta struct {
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Hide        bool   `json:"hide"`
	CustomOrder int    `json:"custom_order"`
}
