package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/LinkedServicesGet.json
func ExampleLinkedServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkedServicesClient().Get(ctx, "mms-eus", "TestLinkWS", "Cluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoperationalinsights.LinkedServicesClientGetResponse{
	// 	LinkedService: armoperationalinsights.LinkedService{
	// 		Name: to.Ptr("TestLinkWS/Cluster"),
	// 		Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testlinkws/linkedservices/cluster"),
	// 		Properties: &armoperationalinsights.LinkedServiceProperties{
	// 			ProvisioningState: to.Ptr(armoperationalinsights.LinkedServiceEntityStatusSucceeded),
	// 			WriteAccessResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster"),
	// 		},
	// 	},
	// }
}
