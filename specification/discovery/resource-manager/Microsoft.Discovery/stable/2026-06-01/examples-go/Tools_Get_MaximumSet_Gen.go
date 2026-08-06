package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/Tools_Get_MaximumSet_Gen.json
func ExampleToolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewToolsClient().Get(ctx, "rgdiscovery", "1ba7068ab4d3671156", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdiscovery.ToolsClientGetResponse{
	// 	Tool: armdiscovery.Tool{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/tools/1ba7068ab4d3671156"),
	// 		Name: to.Ptr("1ba7068ab4d3671156"),
	// 		Tags: map[string]*string{
	// 			"key3848": to.Ptr("um"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/tools"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.ToolProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			Version: to.Ptr("qccygbwif"),
	// 			EnvironmentVariables: map[string]*string{
	// 				"key777": to.Ptr("iyamvfvbaxepw"),
	// 			},
	// 			DefinitionContent: map[string]any{
	// 				"tool_id": "discovery-m1",
	// 				"name": "discovery",
	// 				"description": "Advanced DFT computational tools for molecular geometry optimization and property calculations",
	// 				"actions": []any{
	// 					map[string]any{
	// 						"name": "GeometryOptimization",
	// 						"description": "Optimize geometry of 'xyz's from the input data asset. This is a prerequisite for all other discovery computations.",
	// 						"input_schema": map[string]any{
	// 							"type": "object",
	// 							"properties": map[string]any{
	// 								"inputDataAssetId": map[string]any{
	// 									"type": "string",
	// 									"description": "Identifier of the input data asset",
	// 								},
	// 								"xyzColumnName": map[string]any{
	// 									"type": "string",
	// 									"description": "Column containing xyz data within the input data table asset",
	// 								},
	// 								"outputDataAssetId": map[string]any{
	// 									"type": "string",
	// 									"description": "Identifier to use for the new output data asset which will be created.",
	// 								},
	// 								"basisSet": map[string]any{
	// 									"type": "string",
	// 									"description": "Basis set. Must be one of the supported basis sets (e.g., def2-svp, def2-tzvp).",
	// 								},
	// 							},
	// 							"required": []any{
	// 								"inputDataAssetId",
	// 								"xyzColumnName",
	// 							},
	// 						},
	// 						"command": "python3 submit_dft.py ",
	// 						"environment_variables": []any{
	// 							map[string]any{
	// 								"name": "OUTPUT_DIRECTORY_PATH",
	// 								"value": "{{ outputDataAssetId }}",
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
