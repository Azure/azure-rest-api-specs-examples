package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsPatches_ListByContainerApp.json
func ExampleContainerAppsPatchesClient_NewListByContainerAppPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsPatchesClient().NewListByContainerAppPager("rg", "test-app", &armappcontainers.ContainerAppsPatchesClientListByContainerAppOptions{Filter: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PatchCollection = armappcontainers.PatchCollection{
		// 	Value: []*armappcontainers.ContainerAppsPatchResource{
		// 		{
		// 			Name: to.Ptr("testPatch-25fe4b"),
		// 			Type: to.Ptr("Microsoft.App/containerApps/patches"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app/patches/testPatch-25fe4b"),
		// 			Properties: &armappcontainers.PatchProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-10T12:06:20.342Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-10T12:06:20.342Z"); return t}()),
		// 				PatchApplyStatus: to.Ptr(armappcontainers.PatchApplyStatusNotStarted),
		// 				PatchDetails: []*armappcontainers.PatchDetails{
		// 					{
		// 						DetectionStatus: to.Ptr(armappcontainers.DetectionStatusSucceeded),
		// 						LastDetectionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-10T12:06:19.524Z"); return t}()),
		// 						NewImageName: to.Ptr("testregistry.azurecr.io/test-image:release-1-patched-202210101206185241"),
		// 						NewLayer: &armappcontainers.PatchDetailsNewLayer{
		// 							Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.9-cbl-mariner2.0"),
		// 							FrameworkAndVersion: to.Ptr("dotnet:7.0.9"),
		// 							OSAndVersion: to.Ptr("cbl-mariner2.0"),
		// 						},
		// 						OldLayer: &armappcontainers.PatchDetailsOldLayer{
		// 							Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.7-cbl-mariner2.0"),
		// 							FrameworkAndVersion: to.Ptr("dotnet:7.0.7"),
		// 							OSAndVersion: to.Ptr("cbl-mariner2.0"),
		// 						},
		// 						PatchType: to.Ptr(armappcontainers.PatchTypeFrameworkSecurity),
		// 						TargetContainerName: to.Ptr("test-container"),
		// 						TargetImage: to.Ptr("testregistry.azurecr.io/test-image:release-1-patched-202209101206203421"),
		// 				}},
		// 				TargetContainerAppID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app"),
		// 				TargetEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/test-env"),
		// 				TargetRevisionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/apps/test-app/revisions/test-app--jm3vvry"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testPatch-27c3d5"),
		// 			Type: to.Ptr("Microsoft.App/containerApps/patches"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app/patches/testPatch-27c3d5"),
		// 			Properties: &armappcontainers.PatchProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-10T12:06:20.342Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-20T12:06:20.342Z"); return t}()),
		// 				PatchApplyStatus: to.Ptr(armappcontainers.PatchApplyStatusSucceeded),
		// 				PatchDetails: []*armappcontainers.PatchDetails{
		// 					{
		// 						DetectionStatus: to.Ptr(armappcontainers.DetectionStatusSucceeded),
		// 						LastDetectionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-21T12:06:19.524Z"); return t}()),
		// 						NewImageName: to.Ptr("testregistry.azurecr.io/test-image:release-1-patched-202209101206203421"),
		// 						NewLayer: &armappcontainers.PatchDetailsNewLayer{
		// 							Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.7-cbl-mariner2.0"),
		// 							FrameworkAndVersion: to.Ptr("dotnet:7.0.7"),
		// 							OSAndVersion: to.Ptr("cbl-mariner2.0"),
		// 						},
		// 						OldLayer: &armappcontainers.PatchDetailsOldLayer{
		// 							Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.5-cbl-mariner2.0"),
		// 							FrameworkAndVersion: to.Ptr("dotnet:7.0.5"),
		// 							OSAndVersion: to.Ptr("cbl-mariner2.0"),
		// 						},
		// 						PatchType: to.Ptr(armappcontainers.PatchTypeFrameworkSecurity),
		// 						TargetContainerName: to.Ptr("test-container"),
		// 						TargetImage: to.Ptr("testregistry.azurecr.io/test-image:release-1"),
		// 					},
		// 					{
		// 						DetectionStatus: to.Ptr(armappcontainers.DetectionStatusSucceeded),
		// 						LastDetectionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-21T12:06:19.524Z"); return t}()),
		// 						NewImageName: to.Ptr("testregistry.azurecr.io/test-image:release-2-patched-202209101206203421"),
		// 						NewLayer: &armappcontainers.PatchDetailsNewLayer{
		// 							Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.7-cbl-mariner2.0"),
		// 							FrameworkAndVersion: to.Ptr("dotnet:7.0.7"),
		// 							OSAndVersion: to.Ptr("cbl-mariner2.0"),
		// 						},
		// 						OldLayer: &armappcontainers.PatchDetailsOldLayer{
		// 							Name: to.Ptr("mcr.microsoft.com/dotnet/aspnet:7.0.0-cbl-mariner2.0"),
		// 							FrameworkAndVersion: to.Ptr("dotnet:7.0.0"),
		// 							OSAndVersion: to.Ptr("cbl-mariner2.0"),
		// 						},
		// 						PatchType: to.Ptr(armappcontainers.PatchTypeFrameworkSecurity),
		// 						TargetContainerName: to.Ptr("test-container-2"),
		// 						TargetImage: to.Ptr("testregistry.azurecr.io/test-image:release-2"),
		// 				}},
		// 				TargetContainerAppID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app"),
		// 				TargetEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/test-env"),
		// 				TargetRevisionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerApps/test-app2/revisions/test-app--z79h4oc"),
		// 			},
		// 	}},
		// }
	}
}
