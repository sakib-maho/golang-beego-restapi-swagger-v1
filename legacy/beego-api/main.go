// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package main

import (
	_ "beego-api/routers"

	beego "github.com/beego/beego/v2/server/web"
	
)


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	beego.Run()

}


