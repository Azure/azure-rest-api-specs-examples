package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/ListAppServicePlansByResourceGroup.json
func ExamplePlansClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPlansClient().NewListByResourceGroupPager("testrg123", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PlanCollection = armappservice.PlanCollection{
		// 	Value: []*armappservice.Plan{
		// 		{
		// 			Name: to.Ptr("testsf6141"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.PlanProperties{
		// 				GeoRegion: to.Ptr("East US"),
		// 				IsSpot: to.Ptr(false),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](20),
		// 				NumberOfSites: to.Ptr[int32](4),
		// 				NumberOfWorkers: to.Ptr[int32](19),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				TargetWorkerCount: to.Ptr[int32](0),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("P"),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr("Premium"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testsf7252"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf7252"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.PlanProperties{
		// 				GeoRegion: to.Ptr("East US"),
		// 				IsSpot: to.Ptr(false),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](20),
		// 				NumberOfSites: to.Ptr[int32](4),
		// 				NumberOfWorkers: to.Ptr[int32](19),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				TargetWorkerCount: to.Ptr[int32](0),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("P"),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr("Premium"),
		// 			},
		// 	}},
		// }
	}
}
