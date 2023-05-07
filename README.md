# optionsprobe

## Description
#### take a list of urls from stdin and probe it with options method


## Usage
```
Usage of ./optionsprobe:
  -H string
        headers
  -c int
        max concurrency (default 20)
  -t int
        timeout in millisecond (default 100000)
```
## Example
```
cat urls.txt | optionsprobe
```
### output
```
http://example.com/ OPTIONS
http://example.com/ GET 
http://example.com/ HEAD
http://example.com/ POST
http://example.com/ PUT
```

### control concurrency and timeout
```bash
cat urls.txt | optionsprobe -c 50 -t 5000
```

### tricks 
you can use [assetfinder](https://github.com/tomnomnom/assetfinder) to get subdomains and then probe it with optionsprobe
```bash
assetfinder example.com | optionsprobe
```
or you can use [httprobe](https://github.com/tomnomnom/httprobe/) to get live urls and then probe it with optionsprobe
```bash
cat urls.txt | httprobe | optionsprobe
```

## install
```bash
go install github.com/XORbit01/optionsprobe@latest
```

## build
```bash
git clone https://github.com/XORbit01/optionsprobe.git
cd optionsprobe
go build -o optionsprobe 
./optionsprobe
```

### Binary Release
you can download the binary from release page 
[Releases](https://github.com/XORbit01/optionsprobe/releases/latest)

## Contributing
Pull requests are welcome. please open an issue first to discuss what you would like to change.
you can also contact me on discord:`XORbit#5945`



## Powered By Malwarize
[![image](https://user-images.githubusercontent.com/130087473/232165094-73347c46-71dc-47c0-820a-1eb36657a8c0.png)](https://discord.gg/g9y7D3xCab)


