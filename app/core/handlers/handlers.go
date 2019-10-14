package handlers

import (
	"log"
	"bytes"
	"strings"
	"context"
	"net/http"
	"encoding/json"

	"github.com/cmaddux/string_manipulation/encoding"
	localstrings "github.com/cmaddux/string_manipulation/strings"

	"github.com/gin-gonic/gin"
	es7 "github.com/elastic/go-elasticsearch/v7"
	es7api "github.com/elastic/go-elasticsearch/v7/esapi"
)

func OK() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}

type Text struct {
	Data struct {
		Attributes struct {
			Text string `json:"text" binding:"required"`
		} `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

// GetSpecialCount handler consumes text from request body and
// responds with the count of specila strings in the text.
func GetSpecialCount() gin.HandlerFunc {
	return func (c *gin.Context) {
		var json Text
		if err := c.ShouldBindJSON(&json); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		text := json.Data.Attributes.Text

		rlencoded := encoding.RunLength(text)
		ct := localstrings.CountSpecial(rlencoded)
		c.JSON(
			http.StatusOK,
			gin.H{
				"data": map[string]interface{}{
					"attributes": map[string]interface{}{
						"ct": ct,
					},
				},
			},
		)
		
	}
}

// PostSearch writes a provided search value to elasticserch text
// index.
func PostSearch() gin.HandlerFunc {
	return func (c *gin.Context) {
		es := c.MustGet("elastic").(*es7.Client)	

		var json Text
		if err := c.ShouldBindJSON(&json); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		text := json.Data.Attributes.Text

		var b strings.Builder
		b.WriteString(`{"text" : "`)
		b.WriteString(text)
		b.WriteString(`"}`)

		req := es7api.IndexRequest{
			Index:      "text",
			Body:       strings.NewReader(b.String()),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		defer res.Body.Close()

		if res.IsError() {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// GetSearchValue returns values stored values in text index
// that fuzzy match with (within two transforms of) the
// provided value parameter.
func GetSearchValue() gin.HandlerFunc {
	return func (c *gin.Context) {
		var r  map[string]interface{}

		es := c.MustGet("elastic").(*es7.Client)	

		value := c.Params.ByName("value")

		var buf bytes.Buffer
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"fuzzy": map[string]interface{}{
					"text": value,
				},
			},
		}

		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("text"),
			es.Search.WithBody(&buf),
			es.Search.WithTrackTotalHits(true),
			es.Search.WithPretty(),
		)

		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		defer res.Body.Close()

		if res.IsError() {
			var e map[string]interface{}
			if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
				log.Printf("Unable to parse ES search response body: %s", err)
			} else {
				log.Printf(
					"[%s] %s: %s",
					res.Status(),
					e["error"].(map[string]interface{})["type"],
					e["error"].(map[string]interface{})["reason"],
				)

			}

			c.Status(http.StatusServiceUnavailable)
			return
		}

		if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		var results []map[string]interface{}
		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			data := hit.(map[string]interface{})
			source := data["_source"].(map[string]interface{})
			attributes := map[string]interface{}{
				"text": source["text"],
			}

			item := map[string]interface{}{
				"id": data["_id"],
				"attributes": attributes,
			}

			results = append(results, item)
		}

		if results == nil {
			results = make([]map[string]interface{}, 0)
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"data": results,
			},
		)
	}
}
