package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/SourceControls_Get.json
func ExampleContainerAppsSourceControlsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsSourceControlsClient().Get(ctx, "workerapps-rg-xj", "testcanadacentral", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SourceControl = armappcontainers.SourceControl{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.App/containerapps/sourcecontrols"),
	// 	ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/testcanadacentral/sourcecontrols/current"),
	// 	Properties: &armappcontainers.SourceControlProperties{
	// 		Branch: to.Ptr("master"),
	// 		GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
	// 			BuildEnvironmentVariables: []*armappcontainers.EnvironmentVariable{
	// 				{
	// 					Name: to.Ptr("foo1"),
	// 					Value: to.Ptr("bar1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("foo2"),
	// 					Value: to.Ptr("bar2"),
	// 			}},
	// 			ContextPath: to.Ptr("./"),
	// 			DockerfilePath: to.Ptr("./Dockerfile"),
	// 			Image: to.Ptr("image/tag"),
	// 			RegistryInfo: &armappcontainers.RegistryInfo{
	// 				RegistryURL: to.Ptr("xwang971reg.azurecr.io"),
	// 				RegistryUserName: to.Ptr("xwang971reg"),
	// 			},
	// 		},
	// 		RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
	// 	},
	// }
}
