# neopixel-http

This interface allows you to interact with neopixel led like https://www.adafruit.com/product/1463 
using only high-level http queries, that running in Docker container. 

## Docker 

### Build

```
docker build -f install/python/Dockerfile -t neopixel-interface .
```

### Run 

```
docker run --rm -it --name neopixel-interface --privileged -p 80:80 neopixel-interface
```

### HTTP API 

#### Set pixels to color with opacity

```POST to http://localhost:80```

```
{
	"opacity": 0.3,
	"pixels": [
		{
			"index": 0,
			"r": 255,
			"g": 0,
			"b": 255
		},
		{
			"index": 2,
			"r": 255,
			"g": 255,
			"b": 0
		},
		{
			"index": 3,
			"r": 0,
			"g": 255,
			"b": 255
		}
	]
}
```

![alt text](./docs/example.jpg)

#### Clear all pixels


```POST to http://localhost:80/clear with empty BODY ```

![alt text](./docs/example2.jpg) 

### GPIO 

![alt text](./docs/example3.jpg) 

### Links

* https://github.com/adafruit/Adafruit_CircuitPython_NeoPixel/blob/9e30278fc9fec09f1a4677271e8be3fa266e03c6/neopixel.py