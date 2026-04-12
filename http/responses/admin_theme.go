package responses

// 主配置结构体，对应整个JSON对象
type AdminTheme struct {
	Nav                  Nav                    `json:"nav"`
	Assets               Assets                 `json:"assets"`
	AppName              string                 `json:"app_name"`
	Locale               string                 `json:"locale"`
	Layout               Layout                 `json:"layout"`
	Logo                 string                 `json:"logo"`
	LoginCaptcha         bool                   `json:"login_captcha"`
	LocaleOptions        []LocaleOption         `json:"locale_options"`
	ShowDevelopmentTools bool                   `json:"show_development_tools"`
	SystemThemeSetting   map[string]interface{} `json:"system_theme_setting"` // 使用interface{}因为值可能为null
	EnabledExtensions    []interface{}          `json:"enabled_extensions"`   // 空数组，使用interface{}兼容任何类型
}

// 导航配置结构体
type Nav struct {
	AppendNav  []interface{} `json:"appendNav"`  // 空数组，使用interface{}兼容任何类型
	PrependNav []interface{} `json:"prependNav"` // 空数组，使用interface{}兼容任何类型
}

// 资源配置结构体
type Assets struct {
	Js      []interface{} `json:"js"`      // 空数组，使用interface{}兼容任何类型
	Css     []interface{} `json:"css"`     // 空数组，使用interface{}兼容任何类型
	Scripts []interface{} `json:"scripts"` // 空数组，使用interface{}兼容任何类型
	Styles  []interface{} `json:"styles"`  // 空数组，使用interface{}兼容任何类型
}

// 布局配置结构体
type Layout struct {
	Title            string            `json:"title"`
	Header           Header            `json:"header"`
	LocaleOptions    map[string]string `json:"locale_options"`     // 键值对结构
	KeepAliveExclude []interface{}     `json:"keep_alive_exclude"` // 空数组
	Footer           string            `json:"footer"`
}

// 头部配置结构体
type Header struct {
	Refresh      bool `json:"refresh"`
	Dark         bool `json:"dark"`
	FullScreen   bool `json:"full_screen"`
	LocaleToggle bool `json:"locale_toggle"`
	ThemeConfig  bool `json:"theme_config"`
}

// 语言选项结构体
type LocaleOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// SetLocaleOptions 设置语言选项列表
// 将布局配置中的语言选项映射转换为设置所需的语言选项数组格式
func (c *AdminTheme) SetLocaleOptions() {
	c.LocaleOptions = []LocaleOption{}
	for k, v := range c.Layout.LocaleOptions {
		c.LocaleOptions = append(c.LocaleOptions, LocaleOption{
			Label: v,
			Value: k,
		})
	}
}
