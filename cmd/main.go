package main

import (
	"io/ioutil"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var p1 = `{
	"11" : {
	   "dpgrs" : {
		  "app1" : {
			 "lastUpdated" : "2019-11-12T07:25:11+0000",
			 "name" : "modbus master",
			 "path" : "/home/moxa/data/client/modbus.json",
			 "deviceIp" : "10.144.48.136",
			 "execute" : "/home/moxa/ha-slave",
			 "id" : "app1",
			 "status" : 1,
			 "dpgStatus" : 2,
			 "config" : "{\"id\":\"5a62b848-09c0-11ea-93f0-0090e870a974\",\"name\":\"ha-slave\",\"status\":2,\"workingDirectory\":\"/home/moxa/data/client/configuration.json\",\"configuration\":\"{\\n    \\\"modbus\\\":{\\n        \\\"host\\\":\\\"10.144.49.163\\\",\\n        \\\"port\\\":502,\\n        \\\"deviceId\\\":1,\\n        \\\"intervalSec\\\":1\\n    },\\n    \\\"tags\\\":[\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag1\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 0,\\n            \\\"qty\\\": 2\\n        },\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag2\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 2,\\n            \\\"qty\\\": 2\\n        }\\n    ],\\n    \\\"mqtt\\\":{\\n        \\\"addr\\\":\\\"10.144.49.102:1883\\\",\\n        \\\"topic\\\": \\\"/tags\\\",\\n        \\\"clientId\\\": \\\"mydev1\\\",\\n        \\\"cleanSession\\\": true,\\n        \\\"qos\\\":0\\n    }\\n}\\n\",\"deviceInfo\":{\"ID\":\"b6175f0c-050c-11ea-ac19-0090e8691e1d\",\"Name\":\"My ThingsPro Edge\",\"Model\":\"UC-8112-LX\",\"Ip\":\"10.144.48.137\"},\"execStart\":\"/home/moxa/ha-slave\",\"restart\":\"always\",\"restartSec\":5,\"lastUpdated\":\"2019-11-18T06:06:54+0000\",\"grpcOpts\":{\"connectPort\":\"10.144.48.139:80\",\"connectTimeout\":3000,\"keepaliveInterval\":1000}}"
		  },
		  "app2" : {
			"lastUpdated" : "2019-11-12T07:25:11+0000",
			"name" : "modbus master",
			"path" : "/home/moxa/data/client/mqtt.json",
			"deviceIp" : "10.144.48.136",
			"execute" : "/home/moxa/fake-mqtt",
			"id" : "app2",
			"status" : 1,
			"dpgStatus" : 2,
			"config" : "{\"id\":\"5a62b848-09c0-11ea-93f0-0090e870a974\",\"name\":\"ha-slave\",\"status\":2,\"workingDirectory\":\"/home/moxa/data/client/configuration.json\",\"configuration\":\"{\\n    \\\"modbus\\\":{\\n        \\\"host\\\":\\\"10.144.49.163\\\",\\n        \\\"port\\\":502,\\n        \\\"deviceId\\\":1,\\n        \\\"intervalSec\\\":1\\n    },\\n    \\\"tags\\\":[\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag1\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 0,\\n            \\\"qty\\\": 2\\n        },\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag2\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 2,\\n            \\\"qty\\\": 2\\n        }\\n    ],\\n    \\\"mqtt\\\":{\\n        \\\"addr\\\":\\\"10.144.49.102:1883\\\",\\n        \\\"topic\\\": \\\"/tags\\\",\\n        \\\"clientId\\\": \\\"mydev1\\\",\\n        \\\"cleanSession\\\": true,\\n        \\\"qos\\\":0\\n    }\\n}\\n\",\"deviceInfo\":{\"ID\":\"b6175f0c-050c-11ea-ac19-0090e8691e1d\",\"Name\":\"My ThingsPro Edge\",\"Model\":\"UC-8112-LX\",\"Ip\":\"10.144.48.137\"},\"execStart\":\"/home/moxa/ha-slave\",\"restart\":\"always\",\"restartSec\":5,\"lastUpdated\":\"2019-11-18T06:06:54+0000\",\"grpcOpts\":{\"connectPort\":\"10.144.48.139:80\",\"connectTimeout\":3000,\"keepaliveInterval\":1000}}"
		 }
	   },
	   "deviceIp" : "10.144.48.136",
	   "deviceName" : "MyEdge",
	   "modelName" : "UC-8112-LX",
	   "deviceId" : "11"
	},
	"22" : {
		"dpgrs" : {
		   "app2" : {
			  "lastUpdated" : "2019-11-12T07:25:11+0000",
			  "name" : "ha-slave",
			  "path" : "/home/moxa/data/client/configuration.json",
			  "deviceIp" : "10.144.48.136",
			  "execute" : "/home/moxa/ha-slave",
			  "id" : "app2",
			  "status" : 1,
			  "dpgStatus" : 2,
			  "config" : "{\"id\":\"5a62b848-09c0-11ea-93f0-0090e870a974\",\"name\":\"ha-slave\",\"status\":2,\"workingDirectory\":\"/home/moxa/data/client/configuration.json\",\"configuration\":\"{\\n    \\\"modbus\\\":{\\n        \\\"host\\\":\\\"10.144.49.163\\\",\\n        \\\"port\\\":502,\\n        \\\"deviceId\\\":1,\\n        \\\"intervalSec\\\":1\\n    },\\n    \\\"tags\\\":[\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag1\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 0,\\n            \\\"qty\\\": 2\\n        },\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag2\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 2,\\n            \\\"qty\\\": 2\\n        }\\n    ],\\n    \\\"mqtt\\\":{\\n        \\\"addr\\\":\\\"10.144.49.102:1883\\\",\\n        \\\"topic\\\": \\\"/tags\\\",\\n        \\\"clientId\\\": \\\"mydev1\\\",\\n        \\\"cleanSession\\\": true,\\n        \\\"qos\\\":0\\n    }\\n}\\n\",\"deviceInfo\":{\"ID\":\"b6175f0c-050c-11ea-ac19-0090e8691e1d\",\"Name\":\"My ThingsPro Edge\",\"Model\":\"UC-8112-LX\",\"Ip\":\"10.144.48.137\"},\"execStart\":\"/home/moxa/ha-slave\",\"restart\":\"always\",\"restartSec\":5,\"lastUpdated\":\"2019-11-18T06:06:54+0000\",\"grpcOpts\":{\"connectPort\":\"10.144.48.139:80\",\"connectTimeout\":3000,\"keepaliveInterval\":1000}}"
		   }
		},
		"deviceIp" : "10.144.48.136",
		"deviceName" : "MyEdge",
		"modelName" : "UC-8112-LX",
		"deviceId" : "22"
	 }
}`

var p2 = `{
	"11" : {
	   "dpgrs" : {
		  "app1" : {
			 "lastUpdated" : "2019-11-12T07:25:11+0000",
			 "name" : "ha-slave",
			 "path" : "/home/moxa/data/client/configuration.json",
			 "deviceIp" : "10.144.48.136",
			 "execute" : "/home/moxa/ha-slave",
			 "id" : "app1",
			 "status" : 2,
			 "dpgStatus" : 1,
			 "config" : "{\"id\":\"5a62b848-09c0-11ea-93f0-0090e870a974\",\"name\":\"ha-slave\",\"status\":2,\"workingDirectory\":\"/home/moxa/data/client/configuration.json\",\"configuration\":\"{\\n    \\\"modbus\\\":{\\n        \\\"host\\\":\\\"10.144.49.163\\\",\\n        \\\"port\\\":502,\\n        \\\"deviceId\\\":1,\\n        \\\"intervalSec\\\":1\\n    },\\n    \\\"tags\\\":[\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag1\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 0,\\n            \\\"qty\\\": 2\\n        },\\n        {\\n            \\\"srcNmae\\\": \\\"dev1\\\",\\n            \\\"tagNmae\\\": \\\"tag2\\\",\\n            \\\"valueType\\\": \\\"int\\\",\\n            \\\"addr\\\": 2,\\n            \\\"qty\\\": 2\\n        }\\n    ],\\n    \\\"mqtt\\\":{\\n        \\\"addr\\\":\\\"10.144.49.102:1883\\\",\\n        \\\"topic\\\": \\\"/tags\\\",\\n        \\\"clientId\\\": \\\"mydev1\\\",\\n        \\\"cleanSession\\\": true,\\n        \\\"qos\\\":0\\n    }\\n}\\n\",\"deviceInfo\":{\"ID\":\"b6175f0c-050c-11ea-ac19-0090e8691e1d\",\"Name\":\"My ThingsPro Edge\",\"Model\":\"UC-8112-LX\",\"Ip\":\"10.144.48.137\"},\"execStart\":\"/home/moxa/ha-slave\",\"restart\":\"always\",\"restartSec\":5,\"lastUpdated\":\"2019-11-18T06:06:54+0000\",\"grpcOpts\":{\"connectPort\":\"10.144.48.139:80\",\"connectTimeout\":3000,\"keepaliveInterval\":1000}}"
		  }
	   },
	   "deviceIp" : "10.144.48.136",
	   "deviceName" : "MyEdge",
	   "modelName" : "UC-8112-LX",
	   "deviceId" : "11"
	}
}`

var p3 = `{
	"11" : {
		"dpgrs" : {
			"app1" : {
				"lastUpdated" : "2019-11-12T07:25:11+0000",
				"name" : "ha-slave",
				"path" : "/home/moxa/data/client/configuration.json",
				"deviceIp" : "10.144.48.136",
				"execute" : "/home/moxa/ha-slave",
				"id" : "app1",
				"status" : 3,
				"enable" : true,
				"dpgStatus" : 2,
				"config" : "eyJpZCI6Ijc0MWJmNjA3LTA1MWQtMTFlYS04MWQ3LTAwOTBlODcwYTk3NCIsIm5hbWUiOiJoYS1zbGF2ZSIsInN0YXR1cyI6MSwid29ya2luZ0RpcmVjdG9yeSI6Ii9ob21lL21veGEvZGF0YS9jbGllbnQvY29uZmlndXJhdGlvbi5qc29uIiwiY29uZmlndXJhdGlvbiI6IntcbiAgICBcIm1vZGJ1c1wiOntcbiAgICAgICAgXCJob3N0XCI6XCIxMC4xNDQuNDkuMTYzXCIsXG4gICAgICAgIFwicG9ydFwiOjUwMixcbiAgICAgICAgXCJkZXZpY2VJZFwiOjEsXG4gICAgICAgIFwiaW50ZXJ2YWxTZWNcIjoxXG4gICAgfSxcbiAgICBcInRhZ3NcIjpbXG4gICAgICAgIHtcbiAgICAgICAgICAgIFwic3JjTm1hZVwiOiBcImRldjFcIixcbiAgICAgICAgICAgIFwidGFnTm1hZVwiOiBcInRhZzFcIixcbiAgICAgICAgICAgIFwidmFsdWVUeXBlXCI6IFwiaW50XCIsXG4gICAgICAgICAgICBcImFkZHJcIjogMCxcbiAgICAgICAgICAgIFwicXR5XCI6IDJcbiAgICAgICAgfSxcbiAgICAgICAge1xuICAgICAgICAgICAgXCJzcmNObWFlXCI6IFwiZGV2MVwiLFxuICAgICAgICAgICAgXCJ0YWdObWFlXCI6IFwidGFnMlwiLFxuICAgICAgICAgICAgXCJ2YWx1ZVR5cGVcIjogXCJpbnRcIixcbiAgICAgICAgICAgIFwiYWRkclwiOiAyLFxuICAgICAgICAgICAgXCJxdHlcIjogMlxuICAgICAgICB9XG4gICAgXSxcbiAgICBcIm1xdHRcIjp7XG4gICAgICAgIFwiYWRkclwiOlwiMTAuMTQ0LjQ5LjEwMjoxODgzXCIsXG4gICAgICAgIFwidG9waWNcIjogXCIvdGFnc1wiLFxuICAgICAgICBcImNsaWVudElkXCI6IFwibXlkZXYxXCIsXG4gICAgICAgIFwiY2xlYW5TZXNzaW9uXCI6IHRydWUsXG4gICAgICAgIFwicW9zXCI6MFxuICAgIH1cbn1cbiIsImRldmljZUluZm8iOnsiSUQiOiI5NjBlOWMyNi0wNGVlLTExZWEtYWQyZC0wMDkwZTg3OTliY2YiLCJOYW1lIjoiVUMtODExMi1MWCIsIklwIjoiMTAuMTQ0LjQ4LjEzNiJ9LCJleGVjU3RhcnQiOiIvaG9tZS9tb3hhL2hhLXNsYXZlIiwicmVzdGFydCI6ImFsd2F5cyIsInJlc3RhcnRTZWMiOjUsImxhc3RVcGRhdGVkIjoiMjAxOS0xMS0xMlQwNzowODo1OCswMDAwIiwiZ3JwY09wdHMiOnsiY29ubmVjdFBvcnQiOiIxMC4xNDQuNDguMTM5OjgwIiwiY29ubmVjdFRpbWVvdXQiOjMwMDAsImtlZXBhbGl2ZUludGVydmFsIjoxMDAwfX0="
			}
		},
		"deviceIp" : "10.144.48.136",
		"deviceName" : "MyEdge",
		"modelName" : "UC-8112-LX",
		"deviceId" : "11"
	},
	"22" : {
		"dpgrs" : {
			"app1" : {
				"lastUpdated" : "2019-11-12T07:25:11+0000",
				"name" : "modbus",
				"path" : "/home/moxa/data/client/configuration.json",
				"deviceIp" : "10.144.48.137",
				"execute" : "/home/moxa/ha-slave",
				"id" : "app1",
				"status" : 2,
				"enable" : true,
				"dpgStatus" : 1,
				"config" : "eyJpZCI6Ijc0MWJmNjA3LTA1MWQtMTFlYS04MWQ3LTAwOTBlODcwYTk3NCIsIm5hbWUiOiJoYS1zbGF2ZSIsInN0YXR1cyI6MSwid29ya2luZ0RpcmVjdG9yeSI6Ii9ob21lL21veGEvZGF0YS9jbGllbnQvY29uZmlndXJhdGlvbi5qc29uIiwiY29uZmlndXJhdGlvbiI6IntcbiAgICBcIm1vZGJ1c1wiOntcbiAgICAgICAgXCJob3N0XCI6XCIxMC4xNDQuNDkuMTYzXCIsXG4gICAgICAgIFwicG9ydFwiOjUwMixcbiAgICAgICAgXCJkZXZpY2VJZFwiOjEsXG4gICAgICAgIFwiaW50ZXJ2YWxTZWNcIjoxXG4gICAgfSxcbiAgICBcInRhZ3NcIjpbXG4gICAgICAgIHtcbiAgICAgICAgICAgIFwic3JjTm1hZVwiOiBcImRldjFcIixcbiAgICAgICAgICAgIFwidGFnTm1hZVwiOiBcInRhZzFcIixcbiAgICAgICAgICAgIFwidmFsdWVUeXBlXCI6IFwiaW50XCIsXG4gICAgICAgICAgICBcImFkZHJcIjogMCxcbiAgICAgICAgICAgIFwicXR5XCI6IDJcbiAgICAgICAgfSxcbiAgICAgICAge1xuICAgICAgICAgICAgXCJzcmNObWFlXCI6IFwiZGV2MVwiLFxuICAgICAgICAgICAgXCJ0YWdObWFlXCI6IFwidGFnMlwiLFxuICAgICAgICAgICAgXCJ2YWx1ZVR5cGVcIjogXCJpbnRcIixcbiAgICAgICAgICAgIFwiYWRkclwiOiAyLFxuICAgICAgICAgICAgXCJxdHlcIjogMlxuICAgICAgICB9XG4gICAgXSxcbiAgICBcIm1xdHRcIjp7XG4gICAgICAgIFwiYWRkclwiOlwiMTAuMTQ0LjQ5LjEwMjoxODgzXCIsXG4gICAgICAgIFwidG9waWNcIjogXCIvdGFnc1wiLFxuICAgICAgICBcImNsaWVudElkXCI6IFwibXlkZXYxXCIsXG4gICAgICAgIFwiY2xlYW5TZXNzaW9uXCI6IHRydWUsXG4gICAgICAgIFwicW9zXCI6MFxuICAgIH1cbn1cbiIsImRldmljZUluZm8iOnsiSUQiOiI5NjBlOWMyNi0wNGVlLTExZWEtYWQyZC0wMDkwZTg3OTliY2YiLCJOYW1lIjoiVUMtODExMi1MWCIsIklwIjoiMTAuMTQ0LjQ4LjEzNiJ9LCJleGVjU3RhcnQiOiIvaG9tZS9tb3hhL2hhLXNsYXZlIiwicmVzdGFydCI6ImFsd2F5cyIsInJlc3RhcnRTZWMiOjUsImxhc3RVcGRhdGVkIjoiMjAxOS0xMS0xMlQwNzowODo1OCswMDAwIiwiZ3JwY09wdHMiOnsiY29ubmVjdFBvcnQiOiIxMC4xNDQuNDguMTM5OjgwIiwiY29ubmVjdFRpbWVvdXQiOjMwMDAsImtlZXBhbGl2ZUludGVydmFsIjoxMDAwfX0="
			},
			"app2" : {
				"lastUpdated" : "2019-11-12T07:25:11+0000",
				"name" : "mqtt",
				"path" : "/home/moxa/data/client/configuration.json",
				"deviceIp" : "10.144.48.137",
				"execute" : "/home/moxa/ha-slave",
				"id" : "app2",
				"status" : 2,
				"enable" : false,
				"dpgStatus" : 1,
				"config" : "eyJpZCI6Ijc0MWJmNjA3LTA1MWQtMTFlYS04MWQ3LTAwOTBlODcwYTk3NCIsIm5hbWUiOiJoYS1zbGF2ZSIsInN0YXR1cyI6MSwid29ya2luZ0RpcmVjdG9yeSI6Ii9ob21lL21veGEvZGF0YS9jbGllbnQvY29uZmlndXJhdGlvbi5qc29uIiwiY29uZmlndXJhdGlvbiI6IntcbiAgICBcIm1vZGJ1c1wiOntcbiAgICAgICAgXCJob3N0XCI6XCIxMC4xNDQuNDkuMTYzXCIsXG4gICAgICAgIFwicG9ydFwiOjUwMixcbiAgICAgICAgXCJkZXZpY2VJZFwiOjEsXG4gICAgICAgIFwiaW50ZXJ2YWxTZWNcIjoxXG4gICAgfSxcbiAgICBcInRhZ3NcIjpbXG4gICAgICAgIHtcbiAgICAgICAgICAgIFwic3JjTm1hZVwiOiBcImRldjFcIixcbiAgICAgICAgICAgIFwidGFnTm1hZVwiOiBcInRhZzFcIixcbiAgICAgICAgICAgIFwidmFsdWVUeXBlXCI6IFwiaW50XCIsXG4gICAgICAgICAgICBcImFkZHJcIjogMCxcbiAgICAgICAgICAgIFwicXR5XCI6IDJcbiAgICAgICAgfSxcbiAgICAgICAge1xuICAgICAgICAgICAgXCJzcmNObWFlXCI6IFwiZGV2MVwiLFxuICAgICAgICAgICAgXCJ0YWdObWFlXCI6IFwidGFnMlwiLFxuICAgICAgICAgICAgXCJ2YWx1ZVR5cGVcIjogXCJpbnRcIixcbiAgICAgICAgICAgIFwiYWRkclwiOiAyLFxuICAgICAgICAgICAgXCJxdHlcIjogMlxuICAgICAgICB9XG4gICAgXSxcbiAgICBcIm1xdHRcIjp7XG4gICAgICAgIFwiYWRkclwiOlwiMTAuMTQ0LjQ5LjEwMjoxODgzXCIsXG4gICAgICAgIFwidG9waWNcIjogXCIvdGFnc1wiLFxuICAgICAgICBcImNsaWVudElkXCI6IFwibXlkZXYxXCIsXG4gICAgICAgIFwiY2xlYW5TZXNzaW9uXCI6IHRydWUsXG4gICAgICAgIFwicW9zXCI6MFxuICAgIH1cbn1cbiIsImRldmljZUluZm8iOnsiSUQiOiI5NjBlOWMyNi0wNGVlLTExZWEtYWQyZC0wMDkwZTg3OTliY2YiLCJOYW1lIjoiVUMtODExMi1MWCIsIklwIjoiMTAuMTQ0LjQ4LjEzNiJ9LCJleGVjU3RhcnQiOiIvaG9tZS9tb3hhL2hhLXNsYXZlIiwicmVzdGFydCI6ImFsd2F5cyIsInJlc3RhcnRTZWMiOjUsImxhc3RVcGRhdGVkIjoiMjAxOS0xMS0xMlQwNzowODo1OCswMDAwIiwiZ3JwY09wdHMiOnsiY29ubmVjdFBvcnQiOiIxMC4xNDQuNDguMTM5OjgwIiwiY29ubmVjdFRpbWVvdXQiOjMwMDAsImtlZXBhbGl2ZUludGVydmFsIjoxMDAwfX0="
			}
		},
		"deviceIp" : "10.144.48.137",
		"deviceName" : "MyEdge",
		"modelName" : "UC-3111",
		"deviceId" : "22"
	}
}`

var count = 0

func getInterfaces(c *gin.Context) {
	/* p1, 1 dev to 2 dev
	count++
	if count > 20 {
		count = 0
	}
	if count > 10 {
		c.String(200, p2)
		return
	}
	*/
	c.String(200, p1)

	// c.String(200, p2)
}

func putInterfaces(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	p1 = string(body)
	c.String(200, p1)
}

func replaceInterfaces(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	p1 = string(body)
	c.String(200, p1)
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")
	r.GET("/slaves", getInterfaces)
	r.PUT("/apply", putInterfaces)
	r.PUT("/replace/:id", replaceInterfaces)
	r.Run(":8080")
}
