package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Services_ListByResourceGroup.json
func ExampleServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListByResourceGroupPager("DmsSdkRg", nil)
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
		// page.ServiceList = armdatamigration.ServiceList{
		// 	Value: []*armdatamigration.Service{
		// 		{
		// 			Name: to.Ptr("DmsSdkService1"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService1"),
		// 			Location: to.Ptr("southcentralus"),
		// 			SKU: &armdatamigration.ServiceSKU{
		// 				Name: to.Ptr("GeneralPurpose_4vCores"),
		// 				Size: to.Ptr("4 vCores"),
		// 				Tier: to.Ptr("General Purpose"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DmsSdkService2"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService2"),
		// 			Location: to.Ptr("southcentralus"),
		// 			SKU: &armdatamigration.ServiceSKU{
		// 				Name: to.Ptr("Basic_2vCores"),
		// 				Size: to.Ptr("2 vCores"),
		// 				Tier: to.Ptr("Basic"),
		// 			},
		// 	}},
		// }
	}
}
