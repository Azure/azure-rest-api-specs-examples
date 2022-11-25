package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/DevToolPortals_CreateOrUpdate.json
func ExampleDevToolPortalsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewDevToolPortalsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "default", armappplatform.DevToolPortalResource{
		Properties: &armappplatform.DevToolPortalProperties{
			Features: &armappplatform.DevToolPortalFeatureSettings{
				ApplicationAccelerator: &armappplatform.DevToolPortalFeatureDetail{
					State: to.Ptr(armappplatform.DevToolPortalFeatureStateEnabled),
				},
				ApplicationLiveView: &armappplatform.DevToolPortalFeatureDetail{
					State: to.Ptr(armappplatform.DevToolPortalFeatureStateEnabled),
				},
			},
			Public: to.Ptr(true),
			SsoProperties: &armappplatform.DevToolPortalSsoProperties{
				ClientID:     to.Ptr("00000000-0000-0000-0000-000000000000"),
				ClientSecret: to.Ptr("xxxxx"),
				MetadataURL:  to.Ptr("https://login.microsoft.com/00000000-0000-0000-0000-000000000000/v2.0/.well-known/openid-configuration"),
				Scopes: []*string{
					to.Ptr("openid")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
