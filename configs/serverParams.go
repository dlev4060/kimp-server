package configs

type Configs struct {
	Adress string
}

func InitServerParams() Configs {
	var cfg = Configs{
		Adress: ":9090"}

	return cfg
}
