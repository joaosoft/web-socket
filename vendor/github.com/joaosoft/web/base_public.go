package web

import "strings"

func (r *Base) SetHeader(name string, header []string) {
	r.Headers[strings.Title(name)] = header
}

func (r *Base) GetHeader(name string) string {
	if header, ok := r.Headers[strings.Title(name)]; ok {
		return header[0]
	}
	return ""
}

func (r *Base) SetContentType(contentType ContentType) {
	r.ContentType = contentType
}

func (r *Base) GetContentType() *ContentType {
	if value, ok := r.Headers[string(HeaderContentType)]; ok {
		contentType := ContentType(value[0])
		return &contentType
	}
	return nil
}

func (r *Base) SetCharset(charset Charset) {
	r.Charset = charset
}

func (r *Base) GetCharset() Charset {
	return r.Charset
}

func (r *Base) SetCookie(name string, cookie Cookie) {
	r.Cookies[name] = cookie
}

func (r *Base) GetCookie(name string) *Cookie {
	if cookie, ok := r.Cookies[name]; ok {
		return &cookie
	}
	return nil
}

func (r *Base) SetParam(name string, param []string) {
	r.Params[name] = param
}

func (r *Base) GetParam(name string) string {
	if param, ok := r.Params[name]; ok {
		return param[0]
	}
	return ""
}

func (r *Base) GetParams(name string) []string {
	if param, ok := r.Params[name]; ok {
		return param
	}
	return nil
}

func (r *Base) SetUrlParam(name string, urlParam []string) {
	r.UrlParams[name] = urlParam
}

func (r *Base) GetUrlParam(name string) string {
	if urlParam, ok := r.UrlParams[name]; ok {
		return urlParam[0]
	}
	return ""
}

func (r *Base) GetUrlParams(name string) []string {
	if urlParam, ok := r.UrlParams[name]; ok {
		return urlParam
	}
	return nil
}