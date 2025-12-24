package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/LinkedServicesCreate.json
func ExampleLinkedServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLinkedServicesClient().BeginCreateOrUpdate(ctx, "mms-eus", "TestLinkWS", "Cluster", armoperationalinsights.LinkedService{
		Properties: &armoperationalinsights.LinkedServiceProperties{
			WriteAccessResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster"),
		},
	}, nil)
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
	// res.LinkedService = armoperationalinsights.LinkedService{
	// 	Name: to.Ptr("TestLinkWS/Cluster"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testlinkws/linkedservices/cluster"),
	// 	Properties: &armoperationalinsights.LinkedServiceProperties{
	// 		ProvisioningState: to.Ptr(armoperationalinsights.LinkedServiceEntityStatusSucceeded),
	// 		WriteAccessResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster"),
	// 	},
	// }
}
