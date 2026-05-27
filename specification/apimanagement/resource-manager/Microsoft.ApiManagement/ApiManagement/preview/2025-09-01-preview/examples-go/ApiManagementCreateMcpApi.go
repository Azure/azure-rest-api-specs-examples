package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateMcpApi.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateMcpApi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "mcp-api", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			APIType:     to.Ptr(armapimanagement.APITypeMcp),
			Path:        to.Ptr("mcp-api"),
			Description: to.Ptr("MCP API for AI agent tool discovery and invocation"),
			DisplayName: to.Ptr("MCP API"),
			Protocols: []*armapimanagement.Protocol{
				to.Ptr(armapimanagement.ProtocolHTTPS),
			},
			ServiceURL: to.Ptr("https://mcp-backend.contoso.com"),
			McpProperties: &armapimanagement.McpProperties{
				TransportType: to.Ptr(armapimanagement.McpTransportTypeStreamable),
				Endpoints: []*armapimanagement.McpEndpoint{
					{
						Name:        to.Ptr("message"),
						URITemplate: to.Ptr("/mcp/messages"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIClientCreateOrUpdateResponse{
	// 	APIContract: armapimanagement.APIContract{
	// 		Name: to.Ptr("mcp-api"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/mcp-api"),
	// 		Properties: &armapimanagement.APIContractProperties{
	// 			APIType: to.Ptr(armapimanagement.APITypeMcp),
	// 			Path: to.Ptr("mcp-api"),
	// 			Description: to.Ptr("MCP API for AI agent tool discovery and invocation"),
	// 			APIRevision: to.Ptr("1"),
	// 			DisplayName: to.Ptr("MCP API"),
	// 			IsCurrent: to.Ptr(true),
	// 			IsOnline: to.Ptr(true),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTPS),
	// 			},
	// 			ProvisioningState: to.Ptr("InProgress"),
	// 			ServiceURL: to.Ptr("https://mcp-backend.contoso.com"),
	// 			McpProperties: &armapimanagement.McpProperties{
	// 				TransportType: to.Ptr(armapimanagement.McpTransportTypeStreamable),
	// 				Endpoints: []*armapimanagement.McpEndpoint{
	// 					{
	// 						Name: to.Ptr("message"),
	// 						URITemplate: to.Ptr("/mcp/messages"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
