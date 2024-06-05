package gotaseries

const (
	LocaleFR locale = "fr"
	LocaleEN locale = "en"
	LocaleDE locale = "de"
	LocaleES locale = "es"
	LocaleIT locale = "it"
	LocaleNL locale = "nl"
	LocalePL locale = "pl"
	LocalePT locale = "pt"
)

type locale string

func (l locale) String() string {
	return string(l)
}
