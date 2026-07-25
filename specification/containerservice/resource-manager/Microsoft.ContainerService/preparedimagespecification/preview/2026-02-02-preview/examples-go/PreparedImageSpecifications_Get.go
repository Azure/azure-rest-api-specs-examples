package armcontainerservicepreparedimgspec_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicepreparedimgspec/armcontainerservicepreparedimgspec"
)

// Generated from example definition: 2026-02-02-preview/PreparedImageSpecifications_Get.json
func ExamplePreparedImageSpecificationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicepreparedimgspec.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreparedImageSpecificationsClient().Get(ctx, "rg1", "my-prepared-image-specification", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservicepreparedimgspec.PreparedImageSpecificationsClientGetResponse{
	// 	PreparedImageSpecification: armcontainerservicepreparedimgspec.PreparedImageSpecification{
	// 		Properties: &armcontainerservicepreparedimgspec.PreparedImageSpecificationProperties{
	// 			ContainerImages: []*string{
	// 				to.Ptr("redis:8.0.0"),
	// 			},
	// 			CustomizationScripts: []*armcontainerservicepreparedimgspec.PreparedImageSpecificationScript{
	// 				{
	// 					Name: to.Ptr("initialize-node"),
	// 					ExecutionPoint: to.Ptr(armcontainerservicepreparedimgspec.ExecutionPointNodeImageBuildTime),
	// 					ScriptType: to.Ptr(armcontainerservicepreparedimgspec.ScriptTypeBash),
	// 					Script: to.Ptr("echo \"test prepared image specification\" > /var/log/test-node-customization.txt"),
	// 					PostScriptAction: to.Ptr(armcontainerservicepreparedimgspec.PostScriptActionNone),
	// 				},
	// 			},
	// 			IdentityProfile: &armcontainerservicepreparedimgspec.PreparedImageSpecificationManagedIdentityProfile{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 				ObjectID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
	// 				ClientID: to.Ptr("df4316d4-0ef3-45a3-99d0-2d13a6543aff"),
	// 			},
	// 			Version: to.Ptr("20250101-abcd1234"),
	// 			ProvisioningState: to.Ptr(armcontainerservicepreparedimgspec.ProvisioningStateSucceeded),
	// 		},
	// 		ETag: to.Ptr("ETagValue"),
	// 		Tags: map[string]*string{
	// 			"team": to.Ptr("blue"),
	// 		},
	// 		Location: to.Ptr("westus2"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/preparedImageSpecifications/my-prepared-image-specification"),
	// 		Name: to.Ptr("my-prepared-image-specification"),
	// 		Type: to.Ptr("Microsoft.ContainerService/preparedImageSpecifications"),
	// 		SystemData: &armcontainerservicepreparedimgspec.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicepreparedimgspec.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-02T04:08:43.702Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicepreparedimgspec.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-02T04:08:43.702Z"); return t}()),
	// 		},
	// 	},
	// }
}
