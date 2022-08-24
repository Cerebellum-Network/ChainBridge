module github.com/Cerebellum-Network/ChainBridge

go 1.15

require (
	github.com/Cerebellum-Network/chainbridge-substrate-events v0.0.0-20210506100715-541be1ab4476
	github.com/Cerebellum-Network/chainbridge-utils v1.1.0-cere
	github.com/ChainSafe/log15 v1.0.0
	github.com/centrifuge/go-substrate-rpc-client/v2 v2.0.1
	github.com/ethereum/go-ethereum v1.10.23
	github.com/prometheus/client_golang v1.4.1
	github.com/stretchr/testify v1.7.2
	github.com/urfave/cli/v2 v2.10.2
)

replace github.com/centrifuge/go-substrate-rpc-client/v2 v2.0.1 => github.com/Cerebellum-Network/go-substrate-rpc-client/v2 v2.0.2-cere
