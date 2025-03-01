package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/SharedPrivateLinkResources_Create_MaximumSet_Gen.json
func ExampleSharedPrivateLinkResourcesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSharedPrivateLinkResourcesClient().BeginCreate(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", armdatabasewatcher.SharedPrivateLinkResource{
		Properties: &armdatabasewatcher.SharedPrivateLinkResourceProperties{
			PrivateLinkResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.KeyVault/vaults/kvmo3ej9ih"),
			GroupID:               to.Ptr("vault"),
			RequestMessage:        to.Ptr("request message"),
			DNSZone:               to.Ptr("ec3ae9d410ba"),
			Status:                to.Ptr(armdatabasewatcher.SharedPrivateLinkResourceStatusPending),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.SharedPrivateLinkResourcesClientCreateResponse{
	// 	SharedPrivateLinkResource: &armdatabasewatcher.SharedPrivateLinkResource{
	// 		Properties: &armdatabasewatcher.SharedPrivateLinkResourceProperties{
	// 			PrivateLinkResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.KeyVault/vaults/kvmo3ej9ih"),
	// 			GroupID: to.Ptr("vault"),
	// 			RequestMessage: to.Ptr("request message"),
	// 			DNSZone: to.Ptr("ec3ae9d410ba"),
	// 			Status: to.Ptr(armdatabasewatcher.SharedPrivateLinkResourceStatusPending),
	// 		},
	// 		ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih/sharedPrivateLinkResources/monitoringh22eed"),
	// 		Name: to.Ptr("monitoringh22eed"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers/sharedPrivateLinkResources"),
	// 		SystemData: &armdatabasewatcher.SystemData{
	// 			CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 			CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("mxp"),
	// 			LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		},
	// 	},
	// }
}
