package agollo

import (
	"strings"
	"time"
)

var (
	defaultCluster                    = "default"
	defaultNamespace                  = "application"
	defaultBackupFile                 = ".agollo"
	defaultAutoFetchOnCacheMiss       = false
	defaultFailTolerantOnBackupExists = false
	defaultEnableSLB                  = false
	defaultLongPollInterval           = 1 * time.Second
	defaultEnableHeartBeat            = false
	defaultHeartBeatInterval          = 300 * time.Second
)

type Options struct {
	AppID                      string               // appid
	Cluster                    string               // 默认的集群名称，默认：default
	NameSpaces                 []string             // 预加载命名空间，默认：为空
	ApolloClient               ApolloClient         // apollo HTTP api实现
	Logger                     Logger               // 日志实现类，可以设置自定义实现或者通过NewLogger()创建并设置有效的io.Writer，默认: ioutil.Discard
	AutoFetchOnCacheMiss       bool                 // 自动获取非预设以外的Namespace的配置，默认：false
	LongPollerInterval         time.Duration        // 轮训间隔时间，默认：1s
	BackupFile                 string               // 备份文件存放地址，默认：.agollo
	FailTolerantOnBackupExists bool                 // 服务器连接失败时允许读取备份，默认：false
	Balancer                   Balancer             // ConfigServer负载均衡
	EnableSLB                  bool                 // 启用ConfigServer负载均衡
	RefreshIntervalInSecond    time.Duration        // ConfigServer刷新间隔
	ClientOptions              []ApolloClientOption // 设置apollo HTTP api的配置项
	EnableHeartBeat            bool                 // 是否允许兜底检查，默认：false
	HeartBeatInterval          time.Duration        // 兜底检查间隔时间，默认：300s
}

func getConfigServers(configServerURL string) []string {
	var urls []string
	for _, url := range strings.Split(configServerURL, ",") {
		urls = append(urls, normalizeURL(strings.TrimSpace(url)))
	}

	return urls
}

type Option func(*Options)

func Cluster(cluster string) Option {
	return func(o *Options) {
		o.Cluster = cluster
	}
}

func WithNameSpaces(namespaces []string) Option {
	return func(o *Options) {
		o.NameSpaces = namespaces
	}
}

func WithApolloClient(c ApolloClient) Option {
	return func(o *Options) {
		o.ApolloClient = c
	}
}

func WithLogger(l Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

func AutoFetchOnCacheMiss() Option {
	return func(o *Options) {
		o.AutoFetchOnCacheMiss = true
	}
}

func LongPollerInterval(i time.Duration) Option {
	return func(o *Options) {
		o.LongPollerInterval = i
	}
}

func EnableHeartBeat(b bool) Option {
	return func(o *Options) {
		o.EnableHeartBeat = b
	}
}

func HeartBeatInterval(i time.Duration) Option {
	return func(o *Options) {
		o.HeartBeatInterval = i
	}
}

func BackupFile(backupFile string) Option {
	return func(o *Options) {
		o.BackupFile = backupFile
	}
}

func FailTolerantOnBackupExists() Option {
	return func(o *Options) {
		o.FailTolerantOnBackupExists = true
	}
}

func EnableSLB(b bool) Option {
	return func(o *Options) {
		o.EnableSLB = b
	}
}

func WithBalancer(b Balancer) Option {
	return func(o *Options) {
		o.Balancer = b
	}
}

func ConfigServerRefreshIntervalInSecond(refreshIntervalInSecond time.Duration) Option {
	return func(o *Options) {
		o.RefreshIntervalInSecond = refreshIntervalInSecond
	}
}

func AccessKey(accessKey string) Option {
	return func(o *Options) {
		o.ClientOptions = append(o.ClientOptions, WithAccessKey(accessKey))
	}
}

func WithClientOptions(opts ...ApolloClientOption) Option {
	return func(o *Options) {
		o.ClientOptions = append(o.ClientOptions, opts...)
	}
}

type GetOptions struct {
	// Get时，如果key不存在将返回此值
	DefaultValue string

	// Get时，显示的指定需要获取那个Namespace中的key。非空情况下，优先级顺序为：
	// GetOptions.Namespace > Options.DefaultNamespace > "application"
	Namespace string
}

func (o Options) newGetOptions(opts ...GetOption) GetOptions {
	var getOpts GetOptions
	for _, opt := range opts {
		opt(&getOpts)
	}

	if getOpts.Namespace == "" {
		getOpts.Namespace = "application"
	}

	return getOpts
}

type GetOption func(*GetOptions)

func WithDefault(defVal string) GetOption {
	return func(o *GetOptions) {
		o.DefaultValue = defVal
	}
}

func WithNamespace(namespace string) GetOption {
	return func(o *GetOptions) {
		o.Namespace = namespace
	}
}
