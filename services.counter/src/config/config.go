package config

type AppConfig struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
	I18n struct {
		LocaleDir string   `env:"I18N_LOCALE_DIR"`
		Fallback  string   `env:"I18N_FALLBACK"`
		Locales   []string `env:"I18N_LOCALES"`
	}
}