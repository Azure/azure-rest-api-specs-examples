package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGatewayApi.json
func ExampleGatewayAPIClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayAPIClient().CreateOrUpdate(ctx, "rg1", "apimService1", "gw1", "echo-api", &armapimanagement.GatewayAPIClientCreateOrUpdateOptions{Parameters: &armapimanagement.AssociationContract{
		Properties: &armapimanagement.AssociationContractProperties{
			ProvisioningState: to.Ptr("created"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIContract = armapimanagement.APIContract{
	// 	Name: to.Ptr("echo-api"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/gateways/apis"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1/apis/echo-api"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		APIRevision: to.Ptr("1"),
	// 		IsCurrent: to.Ptr(true),
	// 		SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 			Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 			Query: to.Ptr("subscription-key"),
	// 		},
	// 		Path: to.Ptr(""),
	// 		DisplayName: to.Ptr("EchoApi"),
	// 		Protocols: []*armapimanagement.Protocol{
	// 			to.Ptr(armapimanagement.ProtocolHTTP),
	// 			to.Ptr(armapimanagement.ProtocolHTTPS)},
	// 			ServiceURL: to.Ptr("https://contoso.com/apis/echo"),
	// 		},
	// 	}
}
