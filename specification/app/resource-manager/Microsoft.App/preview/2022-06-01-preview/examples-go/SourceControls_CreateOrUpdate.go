package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/SourceControls_CreateOrUpdate.json
func ExampleContainerAppsSourceControlsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewContainerAppsSourceControlsClient("651f8027-33e8-4ec4-97b4-f6e9f3dc8744", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "workerapps-rg-xj", "testcanadacentral", "current", armappcontainers.SourceControl{
		Properties: &armappcontainers.SourceControlProperties{
			Branch: to.Ptr("master"),
			GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
				AzureCredentials: &armappcontainers.AzureCredentials{
					ClientID:     to.Ptr("<clientid>"),
					ClientSecret: to.Ptr("<clientsecret>"),
					TenantID:     to.Ptr("<tenantid>"),
				},
				ContextPath: to.Ptr("./"),
				Image:       to.Ptr("image/tag"),
				RegistryInfo: &armappcontainers.RegistryInfo{
					RegistryPassword: to.Ptr("<registrypassword>"),
					RegistryURL:      to.Ptr("xwang971reg.azurecr.io"),
					RegistryUserName: to.Ptr("xwang971reg"),
				},
			},
			RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
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
