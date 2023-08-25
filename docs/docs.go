// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/crawler/crawl-chapter": {
            "post": {
                "description": "Crawl chapter from a URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "crawler"
                ],
                "summary": "Crawl Chapter",
                "operationId": "crawl-chapter-details",
                "parameters": [
                    {
                        "description": "Crawl chapter request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CrawlChapterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CrawlChapterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/crawler/crawl-chapters": {
            "post": {
                "description": "Crawl chapters from a URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "crawler"
                ],
                "summary": "Crawl Chapters",
                "operationId": "crawl-chapters",
                "parameters": [
                    {
                        "description": "Crawl chapters request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CrawlChaptersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CrawlChaptersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/crawler/crawl-story": {
            "post": {
                "description": "Crawl a story from a URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "crawler"
                ],
                "summary": "Crawl Story",
                "operationId": "crawl-story",
                "parameters": [
                    {
                        "description": "Crawl story request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.crawlStoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.crawlStoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/stories": {
            "post": {
                "description": "Get all stories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "story"
                ],
                "summary": "Get Stories",
                "operationId": "get-stories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getStoriesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/stories/{sSlug}": {
            "get": {
                "description": "Get Story Details by Slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "story"
                ],
                "summary": "Get Story Details",
                "operationId": "get-story-by-slug",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getStoryBySlugResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/stories/{sSlug}/chapters/{cSlug}": {
            "get": {
                "description": "Get Chapter Details by Slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "story"
                ],
                "summary": "Get Chapter Details",
                "operationId": "get-chapter-by-slug",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.getChapterBySlugResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/translation/do-translate": {
            "post": {
                "description": "Translate a text",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "translation"
                ],
                "summary": "Translate",
                "operationId": "do-translate",
                "parameters": [
                    {
                        "description": "Set up translation",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doTranslateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Translation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/translation/history": {
            "get": {
                "description": "Show all translation history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "translation"
                ],
                "summary": "Show history",
                "operationId": "history",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.historyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Pagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "entity.Author": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Chapter": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "story_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Story": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/entity.Author"
                },
                "author_id": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Category"
                    }
                },
                "chapters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Chapter"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rate": {
                    "$ref": "#/definitions/entity.StoryRate"
                },
                "slug": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                },
                "tid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.StoryRate": {
            "type": "object",
            "properties": {
                "best_value": {
                    "type": "number"
                },
                "count": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "story_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "entity.Translation": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "destination": {
                    "type": "string",
                    "example": "en"
                },
                "id": {
                    "type": "integer"
                },
                "original": {
                    "type": "string",
                    "example": "текст для перевода"
                },
                "source": {
                    "type": "string",
                    "example": "auto"
                },
                "translation": {
                    "type": "string",
                    "example": "text for translation"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "v1.CrawlChapterRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string",
                    "example": "https://truyenfull.vn/tu-cam-270192"
                }
            }
        },
        "v1.CrawlChapterResponse": {
            "type": "object",
            "properties": {
                "chapter": {
                    "$ref": "#/definitions/entity.Chapter"
                }
            }
        },
        "v1.CrawlChaptersRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string",
                    "example": "https://truyenfull.vn/tu-cam-270192"
                }
            }
        },
        "v1.CrawlChaptersResponse": {
            "type": "object"
        },
        "v1.crawlStoryRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string",
                    "example": "https://truyenfull.vn/tu-cam-270192"
                }
            }
        },
        "v1.crawlStoryResponse": {
            "type": "object",
            "properties": {
                "story": {
                    "$ref": "#/definitions/entity.Story"
                }
            }
        },
        "v1.doTranslateRequest": {
            "type": "object",
            "required": [
                "destination",
                "original",
                "source"
            ],
            "properties": {
                "destination": {
                    "type": "string",
                    "example": "en"
                },
                "original": {
                    "type": "string",
                    "example": "текст для перевода"
                },
                "source": {
                    "type": "string",
                    "example": "auto"
                }
            }
        },
        "v1.getChapterBySlugResponse": {
            "type": "object",
            "properties": {
                "chapter": {
                    "$ref": "#/definitions/entity.Chapter"
                }
            }
        },
        "v1.getStoriesResponse": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/common.Pagination"
                },
                "stories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Story"
                    }
                }
            }
        },
        "v1.getStoryBySlugResponse": {
            "type": "object",
            "properties": {
                "story": {
                    "$ref": "#/definitions/entity.Story"
                }
            }
        },
        "v1.historyResponse": {
            "type": "object",
            "properties": {
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Translation"
                    }
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go Clean Template API",
	Description:      "Using a translation service as an example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
