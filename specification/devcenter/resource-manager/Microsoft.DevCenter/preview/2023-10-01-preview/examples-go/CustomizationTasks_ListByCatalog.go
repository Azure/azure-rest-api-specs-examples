package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CustomizationTasks_ListByCatalog.json
func ExampleCustomizationTasksClient_NewListByCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomizationTasksClient().NewListByCatalogPager("rg1", "Contoso", "CentralCatalog", &armdevcenter.CustomizationTasksClientListByCatalogOptions{Top: nil})
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
		// page.CustomizationTaskListResult = armdevcenter.CustomizationTaskListResult{
		// 	Value: []*armdevcenter.CustomizationTask{
		// 		{
		// 			Name: to.Ptr("SampleTask"),
		// 			Type: to.Ptr("Microsoft.DevCenter/devcenters/catalogs/tasks"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/catalogs/CentralCatalog/tasks/SampleTask"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user1"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdevcenter.CustomizationTaskProperties{
		// 				Inputs: map[string]*armdevcenter.CustomizationTaskInput{
		// 					"feed": &armdevcenter.CustomizationTaskInput{
		// 						Type: to.Ptr(armdevcenter.CustomizationTaskInputTypeString),
		// 					},
		// 					"package": &armdevcenter.CustomizationTaskInput{
		// 						Type: to.Ptr(armdevcenter.CustomizationTaskInputTypeString),
		// 						Required: to.Ptr(true),
		// 					},
		// 				},
		// 				Timeout: to.Ptr[int32](30),
		// 			},
		// 	}},
		// }
	}
}
