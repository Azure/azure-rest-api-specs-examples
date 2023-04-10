package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetBuildResult.json
func ExampleBuildServiceClient_GetBuildResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBuildServiceClient().GetBuildResult(ctx, "myResourceGroup", "myservice", "default", "mybuild", "123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BuildResult = armappplatform.BuildResult{
	// 	Name: to.Ptr("123"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/builds/results"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builds/mybuild/results/123"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.BuildResultProperties{
	// 		Name: to.Ptr("123"),
	// 		BuildPodName: to.Ptr("mybuild-default-1"),
	// 		BuildStages: []*armappplatform.BuildStageProperties{
	// 			{
	// 				Name: to.Ptr("prepare"),
	// 				Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateRunning),
	// 			},
	// 			{
	// 				Name: to.Ptr("detect"),
	// 				Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateNotStarted),
	// 			},
	// 			{
	// 				Name: to.Ptr("analyze"),
	// 				Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateNotStarted),
	// 			},
	// 			{
	// 				Name: to.Ptr("restore"),
	// 				Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateNotStarted),
	// 			},
	// 			{
	// 				Name: to.Ptr("build"),
	// 				Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateNotStarted),
	// 			},
	// 			{
	// 				Name: to.Ptr("export"),
	// 				Status: to.Ptr(armappplatform.KPackBuildStageProvisioningStateNotStarted),
	// 		}},
	// 		ProvisioningState: to.Ptr(armappplatform.BuildResultProvisioningStateSucceeded),
	// 	},
	// }
}
