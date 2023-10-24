package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bce66ff64f0e9edc9ea6119d00324058413e81ed/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2023-09-04/examples/SyncIdentityProviders_Get.json
func ExampleSyncIdentityProvidersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSyncIdentityProvidersClient().Get(ctx, "resourceGroup", "resourceName", "childResourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncIdentityProvider = armredhatopenshift.SyncIdentityProvider{
	// 	Name: to.Ptr("mySyncIdentityProvider"),
	// 	Type: to.Ptr("Microsoft.RedHatOpenShift/OpenShiftClusters/SyncIdentityProviders"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.RedHatOpenShift/OpenShiftClusters/resourceName/syncidentityprovider/mySyncIdentityProvider"),
	// 	Properties: &armredhatopenshift.SyncIdentityProviderProperties{
	// 		Resources: to.Ptr("ewogICAgImFwaVZlcnNpb24iOiAiaGl2ZS5vcGVuc2hpZnQuaW8vdjEiLAogICAgImtpbmQiOiAiU3luY0lkZW50aXR5UHJvdmlkZXIiLAogICAgIm1ldGFkYXRhIjogewogICAgICAgICJuYW1lIjogInRlc3QtY2x1c3RlciIsCiAgICAgICAgIm5hbWVzcGFjZSI6ICJhcm8tZjYwYWU4YTItYmNhMS00OTg3LTkwNTYtWFhYWFhYWFhYWFhYIgogICAgfSwKICAgICJzcGVjIjogewogICAgICAgICJjbHVzdGVyRGVwbG95bWVudFJlZnMiOiBbCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogInRlc3QtY2x1c3RlciIKICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgImlkZW50aXR5UHJvdmlkZXJzIjogWwogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAiaHRwYXNzd2QiOiB7CiAgICAgICAgICAgICAgICAgICAgImZpbGVEYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAibmFtZSI6ICJodHBhc3N3ZC1zZWNyZXQiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICJtYXBwaW5nTWV0aG9kIjogImNsYWltIiwKICAgICAgICAgICAgICAgICJuYW1lIjogIkhUUGFzc3dkIiwKICAgICAgICAgICAgICAgICJ0eXBlIjogIkhUUGFzc3dkIgogICAgICAgICAgICB9CiAgICAgICAgXQogICAgfSwKICAgICJzdGF0dXMiOiB7fQp9Cg=="),
	// 	},
	// }
}
