package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ListSwagger.json
func ExampleWorkflowsClient_ListSwagger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowsClient().ListSwagger(ctx, "testResourceGroup", "testWorkflowName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = map[string]any{
	// 	"info":map[string]any{
	// 		"description": "Azure Logic App.",
	// 		"title": "flow",
	// 	},
	// 	"basePath": "/workflows/2e420261710e423490d5d502fe9c4abb/triggers",
	// 	"consumes":[]any{
	// 		"application/json",
	// 	},
	// 	"host": "test-host",
	// 	"paths":map[string]any{
	// 		"/simpleManualTrigger/paths/invoke":map[string]any{
	// 			"post":map[string]any{
	// 				"operationId": "simpleManualTrigger-invoke",
	// 				"description": "Trigger a run of the logic app.",
	// 				"parameters":[]any{
	// 					map[string]any{
	// 						"name": "api-version",
	// 						"type": "string",
	// 						"description": "The service API version.",
	// 						"in": "query",
	// 						"required": true,
	// 					},
	// 					map[string]any{
	// 						"name": "sp",
	// 						"type": "string",
	// 						"description": "The permissions; generally 'read' or 'write'.",
	// 						"in": "query",
	// 						"required": true,
	// 					},
	// 					map[string]any{
	// 						"name": "sv",
	// 						"type": "string",
	// 						"description": "The version number of the query parameters.",
	// 						"in": "query",
	// 						"required": true,
	// 					},
	// 					map[string]any{
	// 						"name": "sig",
	// 						"type": "string",
	// 						"description": "The SHA 256 hash of the entire request URI with an internal key.",
	// 						"in": "query",
	// 						"required": true,
	// 					},
	// 				},
	// 				"responses":map[string]any{
	// 					"default":map[string]any{
	// 						"schema":map[string]any{
	// 							"type": "object",
	// 						},
	// 						"description": "The Logic App Response.",
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	"produces":[]any{
	// 		"application/json",
	// 	},
	// 	"schemes":[]any{
	// 		"http",
	// 	},
	// 	"swagger": "2.0",
	// }
}
