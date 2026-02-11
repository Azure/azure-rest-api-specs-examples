package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/ConfigurationStoresGetSnapshot.json
func ExampleSnapshotsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSnapshotsClient().Get(ctx, "myResourceGroup", "contoso", "mySnapshot", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappconfiguration.SnapshotsClientGetResponse{
	// 	Snapshot: &armappconfiguration.Snapshot{
	// 		Name: to.Ptr("mySnapshot"),
	// 		Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/snapshots"),
	// 		ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/snapshots/mySnapshot"),
	// 		Properties: &armappconfiguration.SnapshotProperties{
	// 			CompositionType: to.Ptr(armappconfiguration.CompositionType("All")),
	// 			Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T22:19:40+00:00"); return t}()),
	// 			Etag: to.Ptr("4f6dd610dd5e4deebc7fbaef685fb903"),
	// 			Filters: []*armappconfiguration.KeyValueFilter{
	// 				{
	// 					Key: to.Ptr("app1/*"),
	// 					Label: to.Ptr("Production"),
	// 				},
	// 			},
	// 			ItemsCount: to.Ptr[int64](71),
	// 			ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
	// 			RetentionPeriod: to.Ptr[int64](3600),
	// 			Size: to.Ptr[int64](100000),
	// 			Tags: map[string]*string{
	// 			},
	// 		},
	// 	},
	// }
}
