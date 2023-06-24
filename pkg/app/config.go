package app

import "issue-detector/pkg/repository"

type Config struct {
	Host        string `config:"APP_HOST" yaml:"host"`
	Port        string `config:"APP_PORT" yaml:"port"`
	SecretToken string `yaml:"secret_token"`

	HostsDB   repository.Config `yaml:"hosts_db"`
	UsersDB   repository.Config `yaml:"users_db"`
	CheckIpDB repository.Config `yaml:"check_ip_history_db"`
}
