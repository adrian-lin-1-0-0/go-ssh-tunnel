package tunnel

type Config struct {
	User    string `yaml:"username"`
	Pwd     string `yaml:"password"`
	KeyPath string `yaml:"sshKeyPath"`
	SvrAddr string `yaml:"serverAddr"`
	SrcAddr string `yaml:"sourceAddr"`
	DstAddr string `yaml:"destAddr"`
}
