package main

type getoptions struct {
	cluster string
	addr    string
	auth    bool
}

// GetOption represents option of get op
type GetOption func(o *getoptions)

// WithCluster sets cluster of get context
func WithCluster(cluster string) GetOption {
	return func(o *getoptions) {
		o.cluster = cluster
	}
}

// WithAddr sets addr for http request instead get from consul
func WithAddr(addr string) GetOption {
	return func(o *getoptions) {
		o.addr = addr
	}
}

// WithAuth Set the GDPR Certify On.
func WithAuth(auth bool) GetOption {
	return func(o *getoptions) {
		o.auth = auth
	}
}

type BConfigClient struct {
	oo getoptions // 函数式选项配置
}

// NewBConfigClient creates instance of BConfigClient
func NewBConfigClient(opts ...GetOption) *BConfigClient {
	oo := getoptions{cluster: "defaultCluster"}
	for _, op := range opts {
		op(&oo)
	}
	c := &BConfigClient{oo: oo}
	return c
}
