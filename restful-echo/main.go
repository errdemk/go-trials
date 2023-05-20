package main

import (
    "strconv"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type link struct {
    Index           int     `json:"index"`
    LocalAddress    string  `json:"localAddress"`
    RemoteAddress   string  `json:"remoteAddress"`
    State           string  `json:"state"`
}

type stat struct {
    Index           int     `json:"index"`
    PacketCount     int     `json:"packetCount"`
    DataSize        int     `json:"dataSize"`
    AverageRTT      float64 `json:"averageRTT"`
}

var links = []link{
    {Index: 1, LocalAddress: "10.10.10.10:9000", RemoteAddress: "20.20.20.20:9000", State: "Active"},
    {Index: 2, LocalAddress: "10.10.10.10:9001", RemoteAddress: "20.20.20.20:9001", State: "Inactive"},
    {Index: 3, LocalAddress: "10.10.10.10:9002", RemoteAddress: "20.20.20.20:9002", State: "Active"},
}

var stats = []stat{
    {Index: 1, PacketCount: 600, DataSize: 256, AverageRTT: 12.5},
    {Index: 2, PacketCount: 700, DataSize: 512, AverageRTT: 5.2},
    {Index: 3, PacketCount: 800, DataSize: 1024, AverageRTT: 1.3},
}

func getLinks(c echo.Context) error {
    return c.JSON(http.StatusOK, links)
}

func getLinkByIndex(c echo.Context) error {
    index := c.Param("index")
    linkIndex, err := strconv.Atoi(index) // string to int

    if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid link index")
	}

    for _, l := range links {
        if l.Index == linkIndex {
            return c.JSON(http.StatusOK, l)
        }
    }
    return c.JSON(http.StatusNotFound, "Link index does not exist")
}

func getLinkByLocalAddress(c echo.Context) error {
    localAddress := c.Param("localAddress")
    // TODO check address regex

    for _, l := range links {
        if l.LocalAddress == localAddress {
            return c.JSON(http.StatusOK, l)
        }
    }
    return c.JSON(http.StatusNotFound, "Invalid link local address")
}

// curl -X POST -H "Content-Type: application/json" -d "{\"state\": \"Inactive\"}" http://localhost:8081/links/1
func updateLinkState(c echo.Context) error {
	index := c.Param("index")
	linkIndex, err := strconv.Atoi(index) // string to int

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid link index")
	}

    for i, l := range links {
        if l.Index == linkIndex {
            var updateReq struct {
                State string `json:"state"`
            }
            if err := c.Bind(&updateReq); err != nil {
                return c.JSON(http.StatusBadRequest, "Invalid request payload")
            }
	        l.State = updateReq.State
            links[i] = l;
	        return c.JSON(http.StatusOK, l)
        }
    }
    // !!!!! Better error handling forced by Echo
    return c.JSON(http.StatusBadRequest, "Link index does not exist")
}

func getStats(c echo.Context) error {
    return c.JSON(http.StatusOK, stats)
}

func getStatByIndex(c echo.Context) error {
    index := c.Param("index")
    statIndex, err := strconv.Atoi(index) // string to int

    if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid stat index")
	}

    for _, s := range stats {
        if s.Index == statIndex {
            return c.JSON(http.StatusOK, s)
        }
    }
    return c.JSON(http.StatusNotFound, "Stat index does not exist")
}

// curl -X DELETE http://localhost:8081/stats
func resetStats(c echo.Context) error {
    for i, s := range stats {
        s.PacketCount = 0
        s.DataSize = 0
        s.AverageRTT = 0.0
        stats[i] = s;
    }
    return c.JSON(http.StatusOK, stats)
}

// curl -X DELETE http://localhost:8081/stats/1
func resetStatByIndex(c echo.Context) error {
    index := c.Param("index")
    statIndex, err := strconv.Atoi(index) // string to int

    if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid stat index")
	}

    for i, s := range stats {
        if s.Index == statIndex {
            s.PacketCount = 0
            s.DataSize = 0
            s.AverageRTT = 0.0
            stats[i] = s;
            return c.JSON(http.StatusOK, s)
        }

    }
    return c.JSON(http.StatusNotFound, "Stat index does not exist")
}

func main() {
	e := echo.New()
  
	// Middleware
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.GET("/links", getLinks)
    e.GET("/links/index/:index", getLinkByIndex)
    e.GET("/links/localAddress/:localAddress", getLinkByLocalAddress)
    e.POST("/links/:index", updateLinkState)

    e.GET("/stats", getStats)
    e.GET("/stats/:index", getStatByIndex)
    e.DELETE("/stats", resetStats)
    e.DELETE("/stats/:index", resetStatByIndex)

	e.Logger.Fatal(e.Start("localhost:8081"))
}