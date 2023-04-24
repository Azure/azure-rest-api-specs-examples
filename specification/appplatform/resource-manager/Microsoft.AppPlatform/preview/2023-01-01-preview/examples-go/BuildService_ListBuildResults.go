package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/BuildService_ListBuildResults.json
func ExampleBuildServiceClient_NewListBuildResultsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBuildServiceClient().NewListBuildResultsPager("myResourceGroup", "myservice", "default", "mybuild", nil)
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
		// page.BuildResultCollection = armappplatform.BuildResultCollection{
		// 	Value: []*armappplatform.BuildResult{
		// 		{
		// 			Name: to.Ptr("123"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/builds/results"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builds/mybuild/results/123"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.BuildResultProperties{
		// 				Name: to.Ptr("123"),
		// 				BuildPodName: to.Ptr("mybuild-default-1"),
		// 				BuildStages: []*armappplatform.BuildStageProperties{
		// 					{
		// 						Name: to.Ptr("prepare"),
		// 						ExitCode: to.Ptr("0"),
		// 						Reason: to.Ptr("Completed"),
		// 						Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateSucceeded),
		// 					},
		// 					{
		// 						Name: to.Ptr("detect"),
		// 						ExitCode: to.Ptr("0"),
		// 						Reason: to.Ptr("Completed"),
		// 						Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateSucceeded),
		// 					},
		// 					{
		// 						Name: to.Ptr("analyze"),
		// 						ExitCode: to.Ptr("0"),
		// 						Reason: to.Ptr("Completed"),
		// 						Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateSucceeded),
		// 					},
		// 					{
		// 						Name: to.Ptr("restore"),
		// 						ExitCode: to.Ptr("0"),
		// 						Reason: to.Ptr("Completed"),
		// 						Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateSucceeded),
		// 					},
		// 					{
		// 						Name: to.Ptr("build"),
		// 						ExitCode: to.Ptr("51"),
		// 						Reason: to.Ptr("Error"),
		// 						Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateFailed),
		// 					},
		// 					{
		// 						Name: to.Ptr("export"),
		// 						ExitCode: to.Ptr("-1"),
		// 						Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateNotStarted),
		// 				}},
		// 				Error: &armappplatform.Error{
		// 					Code: to.Ptr("51"),
		// 					Message: to.Ptr("Build failed in stage build with reason OOMKilled, please refer to https://aka.ms/buildexitcode"),
		// 				},
		// 				ProvisioningState: to.Ptr(armappplatform.BuildResultProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
