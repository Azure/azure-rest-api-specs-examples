package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsBuilds_Get.json
func ExampleContainerAppsBuildsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsBuildsClient().Get(ctx, "rg", "testCapp", "testBuild", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerAppsBuildResource = armappcontainers.ContainerAppsBuildResource{
	// 	Name: to.Ptr("testBuild"),
	// 	Type: to.Ptr("Microsoft.App/containerApps/builds"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/containerApps/testCapp/builds/testBuild"),
	// 	SystemData: &armappcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample@microsoft.com"),
	// 		CreatedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcontainers.ContainerAppsBuildProperties{
	// 		BuildStatus: to.Ptr(armappcontainers.BuildStatusInProgress),
	// 		Configuration: &armappcontainers.ContainerAppsBuildConfiguration{
	// 			BaseOs: to.Ptr("DebianBullseye"),
	// 			EnvironmentVariables: []*armappcontainers.EnvironmentVariable{
	// 				{
	// 					Name: to.Ptr("foo1"),
	// 					Value: to.Ptr("bar1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("foo2"),
	// 					Value: to.Ptr("bar2"),
	// 			}},
	// 			Platform: to.Ptr("dotnetcore"),
	// 			PlatformVersion: to.Ptr("7.0"),
	// 			PreBuildSteps: []*armappcontainers.PreBuildStep{
	// 				{
	// 					Description: to.Ptr("First pre build step."),
	// 					HTTPGet: &armappcontainers.HTTPGet{
	// 						FileName: to.Ptr("output.txt"),
	// 						Headers: []*string{
	// 							to.Ptr("foo"),
	// 							to.Ptr("bar")},
	// 							URL: to.Ptr("https://microsoft.com"),
	// 						},
	// 						Scripts: []*string{
	// 							to.Ptr("echo 'hello'"),
	// 							to.Ptr("echo 'world'")},
	// 						},
	// 						{
	// 							Description: to.Ptr("Second pre build step."),
	// 							HTTPGet: &armappcontainers.HTTPGet{
	// 								FileName: to.Ptr("output.txt"),
	// 								Headers: []*string{
	// 									to.Ptr("foo")},
	// 									URL: to.Ptr("https://microsoft.com"),
	// 								},
	// 								Scripts: []*string{
	// 									to.Ptr("echo 'hello'"),
	// 									to.Ptr("echo 'again'")},
	// 							}},
	// 						},
	// 						DestinationContainerRegistry: &armappcontainers.ContainerRegistryWithCustomImage{
	// 							Image: to.Ptr("test.azurecr.io/repo:tag"),
	// 							Server: to.Ptr("test.azurecr.io"),
	// 						},
	// 						LogStreamEndpoint: to.Ptr("https://foo.azurecontainerapps.dev/logstream"),
	// 						ProvisioningState: to.Ptr(armappcontainers.BuildProvisioningStateSucceeded),
	// 					},
	// 				}
}
