package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsPatches_Get.json
func ExampleContainerAppsPatchesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsPatchesClient().Get(ctx, "rg", "test-app", "testPatch-25fe4b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerAppsPatchResource = armappcontainers.ContainerAppsPatchResource{
	// 	Name: to.Ptr("testPatch-25fe4b"),
	// 	Type: to.Ptr("Microsoft.App/containerApps/patches"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app/patches/testPatch-25fe4b"),
	// 	Properties: &armappcontainers.PatchProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-10T12:06:20.342Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-10T12:06:20.342Z"); return t}()),
	// 		PatchApplyStatus: to.Ptr(armappcontainers.PatchApplyStatusNotStarted),
	// 		PatchDetails: []*armappcontainers.PatchDetails{
	// 			{
	// 				DetectionStatus: to.Ptr(armappcontainers.DetectionStatusSucceeded),
	// 				LastDetectionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-10T12:06:19.524Z"); return t}()),
	// 				NewImageName: to.Ptr("testregistry.azurecr.io/test-image:latest-patched-202210101206"),
	// 				NewLayer: &armappcontainers.PatchDetailsNewLayer{
	// 					Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.7-cbl-mariner2.0"),
	// 					FrameworkAndVersion: to.Ptr("dotnet:7.0.7"),
	// 					OSAndVersion: to.Ptr("cbl-mariner2.0"),
	// 				},
	// 				OldLayer: &armappcontainers.PatchDetailsOldLayer{
	// 					Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.5-cbl-mariner2.0"),
	// 					FrameworkAndVersion: to.Ptr("dotnet:7.0.5"),
	// 					OSAndVersion: to.Ptr("cbl-mariner2.0"),
	// 				},
	// 				PatchType: to.Ptr(armappcontainers.PatchTypeFrameworkSecurity),
	// 				TargetContainerName: to.Ptr("test-container"),
	// 				TargetImage: to.Ptr("testregistry.azurecr.io/test-image:latest"),
	// 		}},
	// 		TargetContainerAppID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app"),
	// 		TargetEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/test-env"),
	// 		TargetRevisionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/apps/test-app/revisions/test-app--jm3vvry"),
	// 	},
	// }
}
