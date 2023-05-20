package main

import (
    "strconv"
	"net/http"
    "github.com/gin-gonic/gin"
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

func getLinks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, links)
}

func getLinkByIndex(c *gin.Context) {
    index := c.Param("index")
    linkIndex, err := strconv.Atoi(index) // string to int

    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link index"})
		return
	}

    for _, l := range links {
        if l.Index == linkIndex {
            c.IndentedJSON(http.StatusOK, l)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Link index does not exist"})
}

func getLinkByLocalAddress(c *gin.Context) {
    localAddress := c.Param("localAddress")
    // TODO check address regex

    for _, l := range links {
        if l.LocalAddress == localAddress {
            c.IndentedJSON(http.StatusOK, l)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid link local address"})
}

// curl -X POST -H "Content-Type: application/json" -d "{\"state\": \"Inactive\"}" http://localhost:8080/links/1
func updateLinkState(c *gin.Context) {
	index := c.Param("index")
	linkIndex, err := strconv.Atoi(index) // string to int

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link index"})
		return
	}

    for i, l := range links {
        if l.Index == linkIndex {
            var updateReq struct {
                State string `json:"state"`
            }
            if err := c.BindJSON(&updateReq); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
                return
            }
	        l.State = updateReq.State
            links[i] = l;
	        c.JSON(http.StatusOK, l)
            return
        }
    }
}

func getStats(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, stats)
}

func getStatByIndex(c *gin.Context) {
    index := c.Param("index")
    statIndex, err := strconv.Atoi(index) // string to int

    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stat index"})
		return
	}

    for _, s := range stats {
        if s.Index == statIndex {
            c.IndentedJSON(http.StatusOK, s)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Stat index does not exist"})
}

// curl -X DELETE http://localhost:8080/stats
func resetStats(c *gin.Context) {
    for i, s := range stats {
        s.PacketCount = 0
        s.DataSize = 0
        s.AverageRTT = 0.0
        stats[i] = s;
    }
    c.JSON(http.StatusOK, stats)
}

// curl -X DELETE http://localhost:8080/stats/1
func resetStatByIndex(c *gin.Context) {
    index := c.Param("index")
    statIndex, err := strconv.Atoi(index) // string to int

    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stat index"})
		return
	}

    for i, s := range stats {
        if s.Index == statIndex {
            s.PacketCount = 0
            s.DataSize = 0
            s.AverageRTT = 0.0
            stats[i] = s;
            c.JSON(http.StatusOK, s)
            return
        }

    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Stat index does not exist"})
}

func main() {
    router := gin.Default()

    // Middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    router.GET("/links", getLinks)
    router.GET("/links/index/:index", getLinkByIndex)
    router.GET("/links/localAddress/:localAddress", getLinkByLocalAddress)
    router.POST("/links/:index", updateLinkState)

    router.GET("/stats", getStats)
    router.GET("/stats/:index", getStatByIndex)
    router.DELETE("/stats", resetStats)
    router.DELETE("/stats/:index", resetStatByIndex)

    router.Run("localhost:8080")
}