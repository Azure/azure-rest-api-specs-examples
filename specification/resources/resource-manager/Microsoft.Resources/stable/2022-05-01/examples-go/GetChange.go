package armchanges_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-05-01/examples/GetChange.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchanges.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "resourceGroup1", "resourceProvider1", "resourceType1", "resourceName1", "1d58d72f-0719-4a48-9228-b7ea682885bf", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ChangeResourceResult = armchanges.ChangeResourceResult{
	// 	Name: to.Ptr("1d58d72f-0719-4a48-9228-b7ea682885bf"),
	// 	Type: to.Ptr("Microsoft.Resources/changes"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId1/resourceGroups/resourceGroup1/providers/resourceProvider1/resourceType1/resourceName1/Microsoft.Resources/changes/1d58d72f-0719-4a48-9228-b7ea682885bf"),
	// 	Properties: &armchanges.ChangeProperties{
	// 		ChangeAttributes: &armchanges.ChangeAttributes{
	// 			ChangesCount: to.Ptr[int64](2),
	// 			CorrelationID: to.Ptr("88420d5d-8d0e-471f-9115-10d34750c617"),
	// 			NewResourceSnapshotID: to.Ptr("6eac9d0f-63b4-4e7f-97a5-740c73757efb"),
	// 			PreviousResourceSnapshotID: to.Ptr("ed90e35a-1661-42cc-a44c-e27f508005be"),
	// 			Timestamp: to.Ptr("2021-11-19T14:29:09.9210000Z"),
	// 		},
	// 		ChangeType: to.Ptr(armchanges.ChangeTypeUpdate),
	// 		Changes: map[string]*armchanges.ChangeBase{
	// 			"properties.provisioningState": &armchanges.ChangeBase{
	// 				ChangeCategory: to.Ptr(armchanges.ChangeCategorySystem),
	// 				NewValue: to.Ptr("Succeeded"),
	// 				PreviousValue: to.Ptr("Updating"),
	// 				PropertyChangeType: to.Ptr(armchanges.PropertyChangeTypeUpdate),
	// 			},
	// 			"tags.key1": &armchanges.ChangeBase{
	// 				ChangeCategory: to.Ptr(armchanges.ChangeCategoryUser),
	// 				NewValue: to.Ptr("someValue"),
	// 				PreviousValue: to.Ptr("null"),
	// 				PropertyChangeType: to.Ptr(armchanges.PropertyChangeTypeInsert),
	// 			},
	// 		},
	// 		TargetResourceID: to.Ptr("/subscriptions/subscriptionId1/resourceGroups/resourceGroup1/providers/resourceProvider1/resourceType1/resourceName1"),
	// 		TargetResourceType: to.Ptr("resourceProvider1/resourceType1"),
	// 	},
	// }
}
