package main

import (
	"./engine"
	"./zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//engine.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/daxinganling",
	//	ParserFunc: parser.ParseUserInfo,
	//})
}