package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetSiteDeploymentStatusSlot.json
func ExampleWebAppsClient_BeginGetSlotSiteDeploymentStatusSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebAppsClient().BeginGetSlotSiteDeploymentStatusSlot(ctx, "rg", "testSite", "stage", "eacfd68b-3bbd-4ad9-99c5-98614d89c8e5", nil)
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
	// res.CsmDeploymentStatus = armappservice.CsmDeploymentStatus{
	// 	Name: to.Ptr("eacfd68b-3bbd-4ad9-99c5-98614d89c8e5"),
	// 	Type: to.Ptr("Microsoft.Web/sites/slots/deploymentStatus"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/sites/testSite/slots/stage/deploymentStatus/eacfd68b-3bbd-4ad9-99c5-98614d89c8e5"),
	// 	Properties: &armappservice.CsmDeploymentStatusProperties{
	// 		DeploymentID: to.Ptr("eacfd68b-3bbd-4ad9-99c5-98614d89c8e5"),
	// 		FailedInstancesLogs: []*string{
	// 		},
	// 		NumberOfInstancesFailed: to.Ptr[int32](0),
	// 		NumberOfInstancesInProgress: to.Ptr[int32](0),
	// 		NumberOfInstancesSuccessful: to.Ptr[int32](1),
	// 		Status: to.Ptr(armappservice.DeploymentBuildStatusRuntimeSuccessful),
	// 	},
	// }
}
