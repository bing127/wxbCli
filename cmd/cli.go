package cmd

import (
	"github.com/gookit/gcli/v2"
	"runtime"
)

const (
	version = "0.0.1"
	logo = `
                            
 _____ _         _____ _ _ 
| __  |_|___ ___|     | |_|
| __ -| |   | . |   --| | |
|_____|_|_|_|_  |_____|_|_|
            |___|          
`
)


func Cli()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp()
	app.Version = version
	app.Description = "wxb命令行工具"
	app.Logo.Text = logo

	// 小程序命令
	app.Add(MiNiApp())
	//
	app.Add(ConvertRemixIcon())
	app.Run()

}