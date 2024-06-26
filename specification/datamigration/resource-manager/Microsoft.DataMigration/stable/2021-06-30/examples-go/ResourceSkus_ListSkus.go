package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/ResourceSkus_ListSkus.json
func ExampleResourceSKUsClient_NewListSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceSKUsClient().NewListSKUsPager(nil)
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
		// page.ResourceSKUsResult = armdatamigration.ResourceSKUsResult{
		// 	Value: []*armdatamigration.ResourceSKU{
		// 		{
		// 			Name: to.Ptr("PP1"),
		// 			Locations: []*string{
		// 				to.Ptr("East Asia"),
		// 				to.Ptr("Southeast Asia"),
		// 				to.Ptr("Australia East"),
		// 				to.Ptr("Australia Southeast"),
		// 				to.Ptr("Brazil South"),
		// 				to.Ptr("Canada Central"),
		// 				to.Ptr("Canada East"),
		// 				to.Ptr("North Europe"),
		// 				to.Ptr("West Europe"),
		// 				to.Ptr("Central India"),
		// 				to.Ptr("South India"),
		// 				to.Ptr("West India"),
		// 				to.Ptr("Japan East"),
		// 				to.Ptr("Japan West"),
		// 				to.Ptr("Korea South"),
		// 				to.Ptr("Korea Central"),
		// 				to.Ptr("UK South"),
		// 				to.Ptr("UK South 2"),
		// 				to.Ptr("UK North"),
		// 				to.Ptr("UK West"),
		// 				to.Ptr("East US 2 EUAP"),
		// 				to.Ptr("Central US"),
		// 				to.Ptr("East US"),
		// 				to.Ptr("East US 2"),
		// 				to.Ptr("North Central US"),
		// 				to.Ptr("South Central US"),
		// 				to.Ptr("West Central US"),
		// 				to.Ptr("West US"),
		// 				to.Ptr("West US 2")},
		// 				ResourceType: to.Ptr("services"),
		// 				Restrictions: []*armdatamigration.ResourceSKURestrictions{
		// 				},
		// 		}},
		// 	}
	}
}
