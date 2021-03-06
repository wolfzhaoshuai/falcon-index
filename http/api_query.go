package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"falcon-index/index"
	"strconv"
	"strings"
)

func configApiQueryRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/v1/fields", func(c *gin.Context) {
		q := c.DefaultQuery("q", "")
		start := c.DefaultQuery("start", "")
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		fmt.Printf("q:%s, start:%s, limit:%d\n", q, start, limit)
		rt, err := index.SearchField(q, start, limit)
		c.JSON(200, gin.H{
			"value": rt,
		})
	})

	router.GET("/api/v1/fieldvalues/field/:f", func(c *gin.Context) {
		f := c.Param("f")
		q := c.DefaultQuery("q", "")
		start := c.DefaultQuery("start", "")
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		fmt.Printf("f:%s, q:%s, start:%s, limit:%d\n", f, q, start, limit)
		rt, err := index.SearchFieldValue(f, q, start, limit)
		c.JSON(200, gin.H{
			"field": f,
			"value": rt,
		})
	})

	router.GET("/api/v1/fields/term/:t", func(c *gin.Context) {
		t := c.Param("t")
		rt, err := index.QueryFieldByTerm(t)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/v1/fields/terms/:ts", func(c *gin.Context) {
		ts := c.Param("ts")
		terms := strings.Split(ts, ",")
		rt, err := index.QueryFieldByTerms(terms)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/v1/fieldvalues/field/:f/terms/:ts", func(c *gin.Context) {
		f := c.Param("f")
		ts := c.Param("ts")
		terms := strings.Split(ts, ",")
		q := c.DefaultQuery("q", "")

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}

		offset_bucket := c.DefaultQuery("offset_bucket", "")
		offset_pos := c.DefaultQuery("offset_position", "")

		var offset *index.Offset
		if offset_bucket != "" && offset_pos != "" {
			offset = &index.Offset{
				Bucket:   []byte(offset_bucket),
				Position: []byte(offset_pos),
			}
		}

		rt, next_offset, err := index.QueryFieldValueByTerms(terms, offset, limit, f, q)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value":           rt,
				"offset_bucket":   string(next_offset.Bucket),
				"offset_position": string(next_offset.Position),
			})
		}
	})

	router.GET("/api/v1/docs/term/:t", func(c *gin.Context) {
		t := c.Param("t")
		start := c.DefaultQuery("start", "")
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		rt, err := index.QueryDocByTerm(t, []byte(start), limit)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/fuzz/endpoint/:p/:search_flag", func(c *gin.Context) {
		pattern := c.Param("p")
		search_flag, err := strconv.Atoi(c.Param("search_flag"))
		if err != nil {
			search_flag = 1
		}
		rt, err := index.FuzzQueryEndpoint(pattern, search_flag)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/fuzz/metric/:p/:search_flag", func(c *gin.Context) {
		pattern := c.Param("p")
		search_flag, err := strconv.Atoi(c.Param("search_flag"))
		if err != nil {
			search_flag = 1
		}
		rt, err := index.FuzzQueryMetric(pattern, search_flag)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/fuzz/counter/:endpoint_name/:patterns/:limit", func(c *gin.Context) {
		endpoint_name := c.Param("endpoint_name")
		patterns := c.Param("patterns")
		pattern_list := strings.Split(patterns, ",")
		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			limit = 100
		}
		rt, err := index.FuzzQueryEndpointMetric(endpoint_name, pattern_list, limit)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/tags/endpoint/:tags", func(c *gin.Context) {
		tags := c.Param("tags")
		k_vs := strings.Split(tags, ",")
		rt, err := index.QueryTagsEndpoint(k_vs)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value": rt,
			})
		}
	})

	router.GET("/api/v1/docs/terms/:ts", func(c *gin.Context) {
		ts := c.Param("ts")
		terms := strings.Split(ts, ",")

		offset_bucket := c.DefaultQuery("offset_bucket", "")
		offset_pos := c.DefaultQuery("offset_position", "")

		var offset *index.Offset
		if offset_bucket != "" && offset_pos != "" {
			offset = &index.Offset{
				Bucket:   []byte(offset_bucket),
				Position: []byte(offset_pos),
			}
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}

		rt, next_offset, err := index.QueryDocByTerms(terms, offset, limit)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"value":           rt,
				"offset_bucket":   string(next_offset.Bucket),
				"offset_position": string(next_offset.Position),
			})
		}
	})
}
