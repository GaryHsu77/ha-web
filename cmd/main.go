package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var p1 = `{
	"11" : {
	   "dpgrs" : {
		  "app1" : {
			 "lastUpdated" : "2019-11-12T07:25:11+0000",
			 "name" : "ha-slave",
			 "path" : "/home/moxa/data/client/configuration.json",
			 "ipaddr" : "10.144.48.136",
			 "execute" : "/home/moxa/ha-slave",
			 "id" : "app1",
			 "status" : 3,
			 "dpgStatus" : 2,
			 "config" : "eyJpZCI6Ijc0MWJmNjA3LTA1MWQtMTFlYS04MWQ3LTAwOTBlODcwYTk3NCIsIm5hbWUiOiJoYS1zbGF2ZSIsInN0YXR1cyI6MSwid29ya2luZ0RpcmVjdG9yeSI6Ii9ob21lL21veGEvZGF0YS9jbGllbnQvY29uZmlndXJhdGlvbi5qc29uIiwiY29uZmlndXJhdGlvbiI6IntcbiAgICBcIm1vZGJ1c1wiOntcbiAgICAgICAgXCJob3N0XCI6XCIxMC4xNDQuNDkuMTYzXCIsXG4gICAgICAgIFwicG9ydFwiOjUwMixcbiAgICAgICAgXCJkZXZpY2VJZFwiOjEsXG4gICAgICAgIFwiaW50ZXJ2YWxTZWNcIjoxXG4gICAgfSxcbiAgICBcInRhZ3NcIjpbXG4gICAgICAgIHtcbiAgICAgICAgICAgIFwic3JjTm1hZVwiOiBcImRldjFcIixcbiAgICAgICAgICAgIFwidGFnTm1hZVwiOiBcInRhZzFcIixcbiAgICAgICAgICAgIFwidmFsdWVUeXBlXCI6IFwiaW50XCIsXG4gICAgICAgICAgICBcImFkZHJcIjogMCxcbiAgICAgICAgICAgIFwicXR5XCI6IDJcbiAgICAgICAgfSxcbiAgICAgICAge1xuICAgICAgICAgICAgXCJzcmNObWFlXCI6IFwiZGV2MVwiLFxuICAgICAgICAgICAgXCJ0YWdObWFlXCI6IFwidGFnMlwiLFxuICAgICAgICAgICAgXCJ2YWx1ZVR5cGVcIjogXCJpbnRcIixcbiAgICAgICAgICAgIFwiYWRkclwiOiAyLFxuICAgICAgICAgICAgXCJxdHlcIjogMlxuICAgICAgICB9XG4gICAgXSxcbiAgICBcIm1xdHRcIjp7XG4gICAgICAgIFwiYWRkclwiOlwiMTAuMTQ0LjQ5LjEwMjoxODgzXCIsXG4gICAgICAgIFwidG9waWNcIjogXCIvdGFnc1wiLFxuICAgICAgICBcImNsaWVudElkXCI6IFwibXlkZXYxXCIsXG4gICAgICAgIFwiY2xlYW5TZXNzaW9uXCI6IHRydWUsXG4gICAgICAgIFwicW9zXCI6MFxuICAgIH1cbn1cbiIsImRldmljZUluZm8iOnsiSUQiOiI5NjBlOWMyNi0wNGVlLTExZWEtYWQyZC0wMDkwZTg3OTliY2YiLCJOYW1lIjoiVUMtODExMi1MWCIsIklwIjoiMTAuMTQ0LjQ4LjEzNiJ9LCJleGVjU3RhcnQiOiIvaG9tZS9tb3hhL2hhLXNsYXZlIiwicmVzdGFydCI6ImFsd2F5cyIsInJlc3RhcnRTZWMiOjUsImxhc3RVcGRhdGVkIjoiMjAxOS0xMS0xMlQwNzowODo1OCswMDAwIiwiZ3JwY09wdHMiOnsiY29ubmVjdFBvcnQiOiIxMC4xNDQuNDguMTM5OjgwIiwiY29ubmVjdFRpbWVvdXQiOjMwMDAsImtlZXBhbGl2ZUludGVydmFsIjoxMDAwfX0="
		  }
	   },
	   "deviceIp" : "10.144.48.136",
	   "deviceName" : "MyEdge",
	   "modelName" : "UC-8112-LX",
	   "deviceId" : "11"
	}
}`

var p2 = `{
	"11" : {
		"dpgrs" : {
			"app1" : {
				"lastUpdated" : "2019-11-12T07:25:11+0000",
				"name" : "ha-slave",
				"path" : "/home/moxa/data/client/configuration.json",
				"ipaddr" : "10.144.48.136",
				"execute" : "/home/moxa/ha-slave",
				"id" : "app1",
				"status" : 3,
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
				"ipaddr" : "10.144.48.137",
				"execute" : "/home/moxa/ha-slave",
				"id" : "app1",
				"status" : 2,
				"dpgStatus" : 1,
				"config" : "eyJpZCI6Ijc0MWJmNjA3LTA1MWQtMTFlYS04MWQ3LTAwOTBlODcwYTk3NCIsIm5hbWUiOiJoYS1zbGF2ZSIsInN0YXR1cyI6MSwid29ya2luZ0RpcmVjdG9yeSI6Ii9ob21lL21veGEvZGF0YS9jbGllbnQvY29uZmlndXJhdGlvbi5qc29uIiwiY29uZmlndXJhdGlvbiI6IntcbiAgICBcIm1vZGJ1c1wiOntcbiAgICAgICAgXCJob3N0XCI6XCIxMC4xNDQuNDkuMTYzXCIsXG4gICAgICAgIFwicG9ydFwiOjUwMixcbiAgICAgICAgXCJkZXZpY2VJZFwiOjEsXG4gICAgICAgIFwiaW50ZXJ2YWxTZWNcIjoxXG4gICAgfSxcbiAgICBcInRhZ3NcIjpbXG4gICAgICAgIHtcbiAgICAgICAgICAgIFwic3JjTm1hZVwiOiBcImRldjFcIixcbiAgICAgICAgICAgIFwidGFnTm1hZVwiOiBcInRhZzFcIixcbiAgICAgICAgICAgIFwidmFsdWVUeXBlXCI6IFwiaW50XCIsXG4gICAgICAgICAgICBcImFkZHJcIjogMCxcbiAgICAgICAgICAgIFwicXR5XCI6IDJcbiAgICAgICAgfSxcbiAgICAgICAge1xuICAgICAgICAgICAgXCJzcmNObWFlXCI6IFwiZGV2MVwiLFxuICAgICAgICAgICAgXCJ0YWdObWFlXCI6IFwidGFnMlwiLFxuICAgICAgICAgICAgXCJ2YWx1ZVR5cGVcIjogXCJpbnRcIixcbiAgICAgICAgICAgIFwiYWRkclwiOiAyLFxuICAgICAgICAgICAgXCJxdHlcIjogMlxuICAgICAgICB9XG4gICAgXSxcbiAgICBcIm1xdHRcIjp7XG4gICAgICAgIFwiYWRkclwiOlwiMTAuMTQ0LjQ5LjEwMjoxODgzXCIsXG4gICAgICAgIFwidG9waWNcIjogXCIvdGFnc1wiLFxuICAgICAgICBcImNsaWVudElkXCI6IFwibXlkZXYxXCIsXG4gICAgICAgIFwiY2xlYW5TZXNzaW9uXCI6IHRydWUsXG4gICAgICAgIFwicW9zXCI6MFxuICAgIH1cbn1cbiIsImRldmljZUluZm8iOnsiSUQiOiI5NjBlOWMyNi0wNGVlLTExZWEtYWQyZC0wMDkwZTg3OTliY2YiLCJOYW1lIjoiVUMtODExMi1MWCIsIklwIjoiMTAuMTQ0LjQ4LjEzNiJ9LCJleGVjU3RhcnQiOiIvaG9tZS9tb3hhL2hhLXNsYXZlIiwicmVzdGFydCI6ImFsd2F5cyIsInJlc3RhcnRTZWMiOjUsImxhc3RVcGRhdGVkIjoiMjAxOS0xMS0xMlQwNzowODo1OCswMDAwIiwiZ3JwY09wdHMiOnsiY29ubmVjdFBvcnQiOiIxMC4xNDQuNDguMTM5OjgwIiwiY29ubmVjdFRpbWVvdXQiOjMwMDAsImtlZXBhbGl2ZUludGVydmFsIjoxMDAwfX0="
			},
			"app2" : {
				"lastUpdated" : "2019-11-12T07:25:11+0000",
				"name" : "mqtt",
				"path" : "/home/moxa/data/client/configuration.json",
				"ipaddr" : "10.144.48.137",
				"execute" : "/home/moxa/ha-slave",
				"id" : "app2",
				"status" : 2,
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
	if count > 10 {
		c.String(200, p2)
		return
	}
	c.String(200, p1)
	*/

	c.String(200, p2)
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")
	r.GET("/slaves", getInterfaces)
	r.Run(":8080")
}
