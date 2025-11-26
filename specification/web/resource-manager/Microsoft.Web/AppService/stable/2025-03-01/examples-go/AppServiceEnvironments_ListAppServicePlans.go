package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_ListAppServicePlans.json
func ExampleEnvironmentsClient_NewListAppServicePlansPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListAppServicePlansPager("test-rg", "test-ase", nil)
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
		// 			Name: to.Ptr("test-asp"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/serverfarms/test-asp"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("Central US EUAP"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappservice.PlanProperties{
		// 				ElasticScaleEnabled: to.Ptr(false),
		// 				GeoRegion: to.Ptr("Central US EUAP"),
		// 				HostingEnvironmentProfile: &armappservice.HostingEnvironmentProfile{
		// 					Name: to.Ptr("test-ase"),
		// 					Type: to.Ptr("Microsoft.Web/hostingEnvironments"),
		// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase"),
		// 				},
		// 				HyperV: to.Ptr(false),
		// 				IsSpot: to.Ptr(false),
		// 				IsXenon: to.Ptr(false),
		// 				MaximumElasticWorkerCount: to.Ptr[int32](0),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](100),
		// 				NumberOfSites: to.Ptr[int32](0),
		// 				PerSiteScaling: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				ResourceGroup: to.Ptr("test-rg"),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				Subscription: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 				TargetWorkerCount: to.Ptr[int32](1),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 				ZoneRedundant: to.Ptr(false),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("I1"),
		// 				Capacity: to.Ptr[int32](0),
		// 				Family: to.Ptr("I"),
		// 				Size: to.Ptr("I1"),
		// 				Tier: to.Ptr("Isolated"),
		// 			},
		// 	}},
		// }
	}
}
