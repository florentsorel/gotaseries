package gotaseries

const (
	LocaleFR LocaleType = "fr"
	LocaleEN LocaleType = "en"
	LocaleDE LocaleType = "de"
	LocaleES LocaleType = "es"
	LocaleIT LocaleType = "it"
	LocaleNL LocaleType = "nl"
	LocalePL LocaleType = "pl"
	LocalePT LocaleType = "pt"
)

type LocaleType string

func (l LocaleType) String() string {
	return string(l)
}
