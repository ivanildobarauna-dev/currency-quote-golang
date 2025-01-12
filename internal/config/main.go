package config

type api_params struct {
	URL string
	ENDPOINT_AVALIABLE_PARITIES string
	ENDPOINT_LAST_COTATION string
	ENDPOINT_HISTORY_COTATION string
	RETRY_TIME_SECONDS int
	RETRY_ATTEMPTS int
}

func GetApiParameters() api_params {

	url:=  "https://economia.awesomeapi.com.br" 

	return api_params{
		URL: url,
		ENDPOINT_AVALIABLE_PARITIES: url + "/json/available",
		ENDPOINT_LAST_COTATION: url + "/last/",
		ENDPOINT_HISTORY_COTATION: url + "/json/daily/",
		RETRY_TIME_SECONDS: 2,
		RETRY_ATTEMPTS: 3,
	}
}
