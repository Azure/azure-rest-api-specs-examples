package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2023-06-01-preview/PrivateLinkScopedResourceUpdatePlatformMetrics.json
func ExamplePrivateLinkScopedResourcesClient_BeginCreateOrUpdate_updateAScopedPlatformMetricsSubscriptionInAPrivateLinkScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkScopedResourcesClient().BeginCreateOrUpdate(ctx, "MyResourceGroup", "MyPrivateLinkScope", "scoped-resource-name", armmonitor.ScopedResource{
		Properties: &armmonitor.ScopedResourceProperties{
			Kind:                 to.Ptr(armmonitor.ScopedResourceKind("PlatformMetrics")),
			LinkedResourceID:     to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444"),
			SubscriptionLocation: to.Ptr("eastus"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.PrivateLinkScopedResourcesClientCreateOrUpdateResponse{
	// 	ScopedResource: armmonitor.ScopedResource{
	// 		Name: to.Ptr("scoped-resource-name"),
	// 		Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/scoped-resource-name"),
	// 		Properties: &armmonitor.ScopedResourceProperties{
	// 			Kind: to.Ptr(armmonitor.ScopedResourceKind("PlatformMetrics")),
	// 			LinkedResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444"),
	// 			ProvisioningState: to.Ptr(armmonitor.ScopedResourceProvisioningStateSucceeded),
	// 			SubscriptionLocation: to.Ptr("eastus"),
	// 		},
	// 		SystemData: &armmonitor.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 			CreatedBy: to.Ptr("bobby@contoso.com"),
	// 			CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("bobby@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
