package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementGetApiTool.json
func ExampleAPIToolClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIToolClient().Get(ctx, "rg1", "apimService1", "github-mcp-api", "findRepositories", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIToolClientGetResponse{
	// 	ToolContract: armapimanagement.ToolContract{
	// 		Name: to.Ptr("findRepositories"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis/tools"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/github-mcp-api/tools/findRepositories"),
	// 		Properties: &armapimanagement.ToolContractProperties{
	// 			OperationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/github-rest-api/operations/findRepositories"),
	// 			Description: to.Ptr("MCP tool to find github repositories by name"),
	// 			DisplayName: to.Ptr("findRepositories"),
	// 		},
	// 	},
	// }
}
