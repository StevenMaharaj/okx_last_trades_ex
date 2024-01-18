# Last Trade streaming example


This program is an example of how to stream the last trade of btc usdt on okex.

## Program structure

- main.go call a stream for okx then listens for trades on a channel that it sent to the streaming.
- stream.go is responisble for the connection to okx. It will pacse the json into the trade struct.
- okxTypes.go contain all the okex type for parsing json in stream.go.
