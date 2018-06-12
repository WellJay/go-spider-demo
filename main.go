package main

import (
	"spider/engine"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/daxinganling",
		ParserFunc: nil,
	})
}
