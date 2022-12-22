package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-09-04/examples/SyncIdentityProviders_Update.json
func ExampleSyncIdentityProvidersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewSyncIdentityProvidersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "resourceGroup", "resourceName", "childResourceName", armredhatopenshift.SyncIdentityProviderUpdate{
		Properties: &armredhatopenshift.SyncIdentityProviderProperties{
			Resources: to.Ptr("ewogICAgImFwaVZlcnNpb24iOiAiaGl2ZS5vcGVuc2hpZnQuaW8vdjEiLAogICAgImtpbmQiOiAiU3luY0lkZW50aXR5UHJvdmlkZXIiLAogICAgIm1ldGFkYXRhIjogewogICAgICAgICJuYW1lIjogInRlc3QtY2x1c3RlciIsCiAgICAgICAgIm5hbWVzcGFjZSI6ICJhcm8tZjYwYWU4YTItYmNhMS00OTg3LTkwNTYtWFhYWFhYWFhYWFhYIgogICAgfSwKICAgICJzcGVjIjogewogICAgICAgICJjbHVzdGVyRGVwbG95bWVudFJlZnMiOiBbCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogInRlc3QtY2x1c3RlciIKICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgImlkZW50aXR5UHJvdmlkZXJzIjogWwogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAiaHRwYXNzd2QiOiB7CiAgICAgICAgICAgICAgICAgICAgImZpbGVEYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAibmFtZSI6ICJodHBhc3N3ZC1zZWNyZXQiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICJtYXBwaW5nTWV0aG9kIjogImNsYWltIiwKICAgICAgICAgICAgICAgICJuYW1lIjogIkhUUGFzc3dkIiwKICAgICAgICAgICAgICAgICJ0eXBlIjogIkhUUGFzc3dkIgogICAgICAgICAgICB9CiAgICAgICAgXQogICAgfSwKICAgICJzdGF0dXMiOiB7fQp9Cg=="),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
