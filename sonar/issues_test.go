package sonar

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalIssuesResponse(t *testing.T) {
	jsonData := `{
		"paging": {
			"pageIndex": 1,
			"pageSize": 100,
			"total": 1
		},
		"issues": [
			{
				"key": "01fc972e-2a3c-433e-bcae-0bd7f88f5123",
				"component": "com.github.kevinsawicki:http-request:com.github.kevinsawicki.http.HttpRequest",
				"project": "com.github.kevinsawicki:http-request",
				"rule": "checkstyle:com.puppycrawl.tools.checkstyle.checks.coding.MagicNumberCheck",
				"issueStatus": "CLOSED",
				"status": "RESOLVED",
				"resolution": "FALSE-POSITIVE",
				"severity": "MINOR",
				"message": "'3' is a magic number.",
				"line": 81,
				"hash": "a227e508d6646b55a086ee11d63b21e9",
				"author": "Developer 1",
				"effort": "2h1min",
				"creationDate": "2013-05-13T17:55:39+0200",
				"updateDate": "2013-05-13T17:55:39+0200",
				"tags": ["bug"],
				"type": "RELIABILITY",
				"comments": [
					{
						"key": "7d7c56f5-7b5a-41b9-87f8-36fa70caa5ba",
						"login": "john.smith",
						"htmlText": "Must be &quot;final&quot;!",
						"markdown": "Must be \"final\"!",
						"updatable": false,
						"createdAt": "2013-05-13T18:08:34+0200"
					}
				],
				"attr": {
					"jira-issue-key": "SONAR-1234"
				},
				"transitions": ["unconfirm", "resolve", "falsepositive"],
				"actions": ["comment"],
				"textRange": {
					"startLine": 2,
					"endLine": 2,
					"startOffset": 0,
					"endOffset": 204
				},
				"flows": [
					{
						"locations": [
							{
								"textRange": {
									"startLine": 16,
									"endLine": 16,
									"startOffset": 0,
									"endOffset": 30
								},
								"msg": "Expected position: 5"
							}
						]
					}
				],
				"ruleDescriptionContextKey": "spring",
				"cleanCodeAttributeCategory": "INTENTIONAL",
				"cleanCodeAttribute": "CLEAR",
				"impacts": [
					{
						"softwareQuality": "MAINTAINABILITY",
						"severity": "HIGH"
					}
				]
			}
		],
		"components": [
			{
				"key": "com.github.kevinsawicki:http-request:src/main/java/com/github/kevinsawicki/http/HttpRequest.java",
				"enabled": true,
				"qualifier": "FIL",
				"name": "HttpRequest.java",
				"longName": "src/main/java/com/github/kevinsawicki/http/HttpRequest.java",
				"path": "src/main/java/com/github/kevinsawicki/http/HttpRequest.java"
			}
		],
		"rules": [
			{
				"key": "checkstyle:com.puppycrawl.tools.checkstyle.checks.coding.MagicNumberCheck",
				"name": "Magic Number",
				"status": "READY",
				"lang": "java",
				"langName": "Java"
			}
		],
		"users": [
			{
				"login": "admin",
				"name": "Administrator",
				"active": true,
				"avatar": "ab0ec6adc38ad44a15105f207394946f"
			}
		]
	}`

	var response IssuesResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	assert.NoError(t, err)

	assert.Equal(t, 1, response.Paging.PageIndex)
	assert.Equal(t, 1, len(response.Issues))
	assert.Equal(t, "01fc972e-2a3c-433e-bcae-0bd7f88f5123", response.Issues[0].Key)
	assert.Equal(t, "com.github.kevinsawicki:http-request:com.github.kevinsawicki.http.HttpRequest", response.Issues[0].Component)
	assert.Equal(t, "Magic Number", response.Rules[0].Name)
	assert.Equal(t, "admin", response.Users[0].Login)
}