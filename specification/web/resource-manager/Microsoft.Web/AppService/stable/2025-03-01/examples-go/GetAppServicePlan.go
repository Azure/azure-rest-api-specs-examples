package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetAppServicePlan.json
func ExamplePlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPlansClient().Get(ctx, "testrg123", "testsf6141", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Plan = armappservice.Plan{
	// 	Name: to.Ptr("testsf6141"),
	// 	Type: to.Ptr("Microsoft.Web/serverfarms"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
	// 	Kind: to.Ptr("app"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.PlanProperties{
	// 		GeoRegion: to.Ptr("East US"),
	// 		IsSpot: to.Ptr(false),
	// 		MaximumNumberOfWorkers: to.Ptr[int32](20),
	// 		NumberOfSites: to.Ptr[int32](4),
	// 		NumberOfWorkers: to.Ptr[int32](19),
	// 		ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 		Reserved: to.Ptr(false),
	// 		Status: to.Ptr(armappservice.StatusOptionsReady),
	// 		TargetWorkerCount: to.Ptr[int32](0),
	// 		TargetWorkerSizeID: to.Ptr[int32](0),
	// 	},
	// 	SKU: &armappservice.SKUDescription{
	// 		Name: to.Ptr("P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("P"),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr("Premium"),
	// 	},
	// }
}
