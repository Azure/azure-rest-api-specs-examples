package armchanges_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-05-01/examples/ListChanges.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchanges.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager("resourceGroup1", "resourceProvider1", "resourceType1", "resourceName1", &armchanges.ClientListOptions{Top: nil,
		SkipToken: nil,
	})
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
		// page.ChangeResourceListResult = armchanges.ChangeResourceListResult{
		// 	Value: []*armchanges.ChangeResourceResult{
		// 		{
		// 			Name: to.Ptr("a9f34285-13a2-e79c-f468-cfb71c7bd227"),
		// 			Type: to.Ptr("Microsoft.Resources/changes"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId1/resourceGroups/resourceGroup1/providers/resourceProvider1/resourceType1/resourceName1/Microsoft.Resources/changes/a9f34285-13a2-e79c-f468-cfb71c7bd227"),
		// 			Properties: &armchanges.ChangeProperties{
		// 				ChangeAttributes: &armchanges.ChangeAttributes{
		// 					ChangesCount: to.Ptr[int64](2),
		// 					CorrelationID: to.Ptr("88420d5d-8d0e-471f-9115-10d34750c617"),
		// 					NewResourceSnapshotID: to.Ptr("6eac9d0f-63b4-4e7f-97a5-740c73757efb"),
		// 					PreviousResourceSnapshotID: to.Ptr("ed90e35a-1661-42cc-a44c-e27f508005be"),
		// 					Timestamp: to.Ptr("2021-11-19T14:29:09.9210000Z"),
		// 				},
		// 				ChangeType: to.Ptr(armchanges.ChangeTypeUpdate),
		// 				Changes: map[string]*armchanges.ChangeBase{
		// 					"properties.provisioningState": &armchanges.ChangeBase{
		// 						ChangeCategory: to.Ptr(armchanges.ChangeCategorySystem),
		// 						NewValue: to.Ptr("Succeeded"),
		// 						PreviousValue: to.Ptr("Updating"),
		// 						PropertyChangeType: to.Ptr(armchanges.PropertyChangeTypeUpdate),
		// 					},
		// 					"tags.key1": &armchanges.ChangeBase{
		// 						ChangeCategory: to.Ptr(armchanges.ChangeCategoryUser),
		// 						NewValue: to.Ptr("someValue"),
		// 						PreviousValue: to.Ptr("null"),
		// 						PropertyChangeType: to.Ptr(armchanges.PropertyChangeTypeInsert),
		// 					},
		// 				},
		// 				TargetResourceID: to.Ptr("/subscriptions/subscriptionId1/resourceGroups/resourceGroup1/providers/resourceProvider1/resourceType1/resourceName1"),
		// 				TargetResourceType: to.Ptr("resourceProvider1/resourceType1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a9f34285-13a2-e79c-f468-cfb71c7bd227"),
		// 			Type: to.Ptr("Microsoft.Resources/changes"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId1/resourceGroups/resourceGroup1/providers/resourceProvider1/resourceType1/resourceName1/Microsoft.Resources/changes/a9f34285-13a2-e79c-f468-cfb71c7bd227"),
		// 			Properties: &armchanges.ChangeProperties{
		// 				ChangeAttributes: &armchanges.ChangeAttributes{
		// 					ChangesCount: to.Ptr[int64](0),
		// 					CorrelationID: to.Ptr("88420d5d-8d0e-471f-9115-10d34750c617"),
		// 					NewResourceSnapshotID: to.Ptr("4db20fc0-de17-4cdd-92d8-fd6bf94b9fd9"),
		// 					PreviousResourceSnapshotID: to.Ptr("b09f5e52-0b46-4d13-84a9-08653d39fed6"),
		// 					Timestamp: to.Ptr("2021-11-19T14:29:09.9210000Z"),
		// 				},
		// 				ChangeType: to.Ptr(armchanges.ChangeTypeCreate),
		// 				Changes: map[string]*armchanges.ChangeBase{
		// 				},
		// 				TargetResourceID: to.Ptr("/subscriptions/subscriptionId1/resourceGroups/resourceGroup1/providers/resourceProvider1/resourceType1/resourceName1"),
		// 				TargetResourceType: to.Ptr("resourceProvider1/resourceType1"),
		// 			},
		// 	}},
		// }
	}
}
