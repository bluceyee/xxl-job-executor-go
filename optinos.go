package xxl

import (
	"github.com/go-basic/ipv4"
	"time"
)

type Options struct {
	ServerAddr   string        `json:"server_addr"`   //调度中心地址
	AccessToken  string        `json:"access_token"`  //请求令牌
	Timeout      time.Duration `json:"timeout"`       //接口超时时间
	ExecutorIp   string        `json:"executor_ip"`   //执行器IP(如果是docker且不在同一个网段，需指定成宿主机ip)
	ExecutorPort string        `json:"executor_port"` //执行器端口(如果是docker且不在同一个网段，需指定成宿主机映射端口)
	LocalIp      string        `json:"local_ip"`      //本地(执行器)IP(可自行获取)
	LocalPort    string        `json:"local_port"`    //本地(执行器)端口
	RegistryKey  string        `json:"registry_key"`  //执行器名称
	LogDir       string        `json:"log_dir"`       //日志目录

	l Logger //日志处理
}

func newOptions(opts ...Option) Options {
	opt := Options{
		ExecutorIp:   ipv4.LocalIP(),
		LocalIp:      ipv4.LocalIP(),
		ExecutorPort: DefaultExecutorPort,
		LocalPort:    DefaultExecutorPort,
		RegistryKey:  DefaultRegistryKey,
	}

	for _, o := range opts {
		o(&opt)
	}

	if opt.l == nil {
		opt.l = &logger{}
	}

	return opt
}

type Option func(o *Options)

var (
	DefaultExecutorPort = "9999"
	DefaultRegistryKey  = "golang-jobs"
)

// ServerAddr 设置调度中心地址
func ServerAddr(addr string) Option {
	return func(o *Options) {
		o.ServerAddr = addr
	}
}

// AccessToken 请求令牌
func AccessToken(token string) Option {
	return func(o *Options) {
		o.AccessToken = token
	}
}

// ExecutorIp 设置执行器IP
func ExecutorIp(ip string) Option {
	return func(o *Options) {
		if len(ip) >0  {
			o.ExecutorIp = ip
		}
	}
}

// ExecutorPort 设置执行器端口
func ExecutorPort(port string) Option {
	return func(o *Options) {
		if len(port) >0 {
			o.ExecutorPort = port
		}
	}
}

// ExecutorIp 设置执行器IP
func LocalIp(ip string) Option {
	return func(o *Options) {
		if len(ip) > 0 {
			o.LocalIp = ip
		}
	}
}

// ExecutorPort 设置执行器端口
func LocalPort(port string) Option {
	return func(o *Options) {
		if len(port) > 0 {
			o.LocalPort = port
		}
	}
}

// RegistryKey 设置执行器标识
func RegistryKey(registryKey string) Option {
	return func(o *Options) {
		o.RegistryKey = registryKey
	}
}

// SetLogger 设置日志处理器
func SetLogger(l Logger) Option {
	return func(o *Options) {
		o.l = l
	}
}
