package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/CreateArcIdentity.json
func ExampleArcSettingsClient_BeginCreateIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewArcSettingsClient().BeginCreateIdentity(ctx, "test-rg", "myCluster", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArcIdentityResponse = armazurestackhci.ArcIdentityResponse{
	// 	Properties: &armazurestackhci.ArcIdentityResponseProperties{
	// 		ArcApplicationClientID: to.Ptr("7b93bf67-60ac-4909-a987-ac438e69f9ba"),
	// 		ArcApplicationObjectID: to.Ptr("400bd05f-395f-45a6-ba75-72601df80107"),
	// 		ArcApplicationTenantID: to.Ptr("bdb2c88c-9cfd-4e19-927d-51e875f6912b"),
	// 		ArcServicePrincipalObjectID: to.Ptr("00cc4014-482e-4de9-9932-83415cc75f45"),
	// 	},
	// }
}
