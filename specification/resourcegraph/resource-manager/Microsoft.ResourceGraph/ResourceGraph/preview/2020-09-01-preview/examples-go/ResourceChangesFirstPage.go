package armresourcegraph_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: 2020-09-01-preview/ResourceChangesFirstPage.json
func ExampleClient_ResourceChanges_firstPageQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ResourceChanges(ctx, armresourcegraph.ResourceChangesRequestParameters{
		Top: to.Ptr[int32](2),
		Interval: &armresourcegraph.ResourceChangesRequestParametersInterval{
			End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T12:09:03.141Z"); return t }()),
			Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-30T12:09:03.141Z"); return t }()),
		},
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourcegraph.ClientResourceChangesResponse{
	// 	ResourceChangeList: armresourcegraph.ResourceChangeList{
	// 		SkipToken: "ew0KICAiJGlkIjogIjEiLA0KICAiRW5kVGltZSI6ICJcL0RhdGUoMTU1MDc0NT",
	// 		Changes: []*armresourcegraph.ResourceChangeData{
	// 			{
	// 				AfterSnapshot: &armresourcegraph.ResourceChangeDataAfterSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T01:54:24.42Z"); return t}()),
	// 				},
	// 				BeforeSnapshot: &armresourcegraph.ResourceChangeDataBeforeSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T01:32:05.993Z"); return t}()),
	// 				},
	// 				ChangeID: to.Ptr("2db0ad2d-f6f0-4f46-b529-5c4e8c494648"),
	// 				ChangeType: to.Ptr(armresourcegraph.ChangeTypeUpdate),
	// 				ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
	// 			},
	// 			{
	// 				AfterSnapshot: &armresourcegraph.ResourceChangeDataAfterSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-30T21:12:31.337Z"); return t}()),
	// 				},
	// 				BeforeSnapshot: &armresourcegraph.ResourceChangeDataBeforeSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-30T10:30:19.68Z"); return t}()),
	// 				},
	// 				ChangeID: to.Ptr("9dc352cb-b7c1-4198-9eda-e5e3ed66aec8"),
	// 				ChangeType: to.Ptr(armresourcegraph.ChangeTypeCreate),
	// 				ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
	// 			},
	// 		},
	// 	},
	// }
}
