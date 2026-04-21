package i18n

import (
	"strings"
)

var locales = map[string]map[string]string{
	"kk": {
		"registration.start":           "✅ Қош келдіңіз! Ботты пайдалану үшін тіркеуден өтіңіз:",
		"registration.ask_name":        "✏️ Аты-жөніңізді еңгізіңіз:",
		"registration.ask_birthday":    "✏️ Туған күніңізді еңгізіңіз:",
		"registration.ask_nationality": "✏️ Ұлтыңызды еңгізіңіз:",
		"registration.ask_workplace":   "✏️ Жұмыс орныңызды еңгізіңіз:",
		"registration.ask_address":     "✏️ Тұрақты мекенжайыңызды еңгізіңіз:",
		"registration.ask_telephone":   "✏️ Байланыс нөміріңізді еңгізіңіз:",
	},
}

func T(lang string, keys ...string) string {
	if l, ok := locales[lang]; ok {
		fullKey := strings.Join(keys, ".")

		if val, ok := l[fullKey]; ok {
			return val
		}
	}

	return strings.Join(keys, ".")
}
