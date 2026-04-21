package i18n

import (
	"fmt"
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
		"test.start":                   "📝 Деңгейді анықтау (диагностикалық тест)\n\nҚысқа диагностикалық тест арқылы қазақ тілі деңгейіңіз анықталады:",
		"test.result":                  "Сіздің деңгейіңіз: %s (%s)",
		"test.A1":                      "Қарапайым",
		"test.A2":                      "Базалық",
		"test.B1":                      "Орта",
		"test.B2":                      "Ортадан жоғары",
		"test.C1":                      "Жоғары",
		"format.start":                 "✅ Сабақ форматын таңдаңыз:",
		"format.online":                "🌐 Онлайн",
		"format.offline":               "📚 Оффлайн",
		"format.location":              "🏫 Мекеме атауы:\n«Ақтөбе облысының Тілдерді дамыту басқармасы» ММ «Тілдерді оқыту орталығы» КММ\n\nМекенжайы:\nАқтөбе қаласы, Тургенев көшесі, 86\n\nБайланыс нөмірі:\n+7 (7132) 46 78 68",
	},
}

func T(lang string, key string, args ...interface{}) string {
	if l, ok := locales[lang]; ok {
		if val, ok := l[key]; ok {
			if len(args) > 0 {
				return fmt.Sprintf(val, args...)
			}
			return val
		}
	}
	return key
}
