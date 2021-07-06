# eth-token-holders
api to find the amount of holders for a particular ethereum or binance-smart-chain token

## usage

```
curl https://eth-token-holders.cloud.rileysnyder.org/<chain>/<contract address>
```

```
curl https://eth-token-holders.cloud.rileysnyder.org/ethereum/0x1f9840a85d5af5bf1d1762f925bdaddc4201f984
249,430

curl https://eth-token-holders.cloud.rileysnyder.org/binance-smart-chain/0x84750e8EfacC62be5AFe7221Cf149A2520Cb1b60
2,773 addresses
```

## runnning

```
./eth-token-holders -addr 0.0.0.0:8080
```
