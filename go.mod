module github.com/Cerebellum-Network/ChainBridge

go 1.15

require (
	github.com/Cerebellum-Network/chainbridge-substrate-events v0.0.0-20210506100715-541be1ab4476
	github.com/Cerebellum-Network/chainbridge-utils v1.1.0-cere
	github.com/ChainSafe/log15 v1.0.0
	github.com/aristanetworks/goarista v0.0.0-20200609010056-95bcf8053598 // indirect
	github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.8
	github.com/ethereum/go-ethereum v1.10.17
	github.com/prometheus/client_golang v1.4.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
)

replace github.com/Cerebellum-Network/chainbridge-substrate-events => ../chainbridge-substrate-events

replace github.com/Cerebellum-Network/chainbridge-utils => ../chainbridge-utils

replace github.com/centrifuge/go-substrate-rpc-client/v4 => ../go-substrate-rpc-client
