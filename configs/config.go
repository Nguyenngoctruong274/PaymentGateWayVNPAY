package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type VNPayConfig struct {
	Vnpay struct {
		ReturnURL   string `yaml:"return_Url"`
		CancelURL   string `yaml:"cancel_Url"`
		VnpVersion  string `yaml:"vnp_Version"`
		VnpCommand  string `yaml:"vnp_Command"`
		VnpTmnCode  string `yaml:"vnp_TmnCode"`
		VnpCurrCode string `yaml:"vnp_CurrCode"`
		VnpUrl      string `yaml:"vnp_Url"`
	} `yaml:"vnpay"`
}

// LoadConfig là một hàm để load cấu hình từ file YAML
func LoadConfig(filename string) (*VNPayConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	var config VNPayConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil, err
	}

	return &config, nil
}
