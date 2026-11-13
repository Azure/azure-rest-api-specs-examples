package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v3"
)

// Generated from example definition: 2026-09-01/application/VMApplications_Get.json
func ExampleVMApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMApplicationsClient().Get(ctx, "myResourceGroup", "myMachine", "myApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridcompute.VMApplicationsClientGetResponse{
	// 	VMApplication: armhybridcompute.VMApplication{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine/applications/myApplication"),
	// 		Name: to.Ptr("myApplication"),
	// 		Type: to.Ptr("Microsoft.HybridCompute/machines/applications"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armhybridcompute.VMApplicationProperties{
	// 			Order: to.Ptr[int32](1),
	// 			PackageReferenceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/galleryResourceGroup/providers/Microsoft.Compute/galleries/myGallery/applications/myApplication/versions/1.0.0"),
	// 			TreatFailureAsDeploymentFailure: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
	// 			InstanceView: &armhybridcompute.VMApplicationInstanceView{
	// 				Name: to.Ptr("myApplication"),
	// 				Version: to.Ptr("1.0.0"),
	// 				Result: to.Ptr("SUCCESS"),
	// 				ActionsPerformed: []*armhybridcompute.VMApplicationInstanceViewAction{
	// 					{
	// 						Operation: to.Ptr("install"),
	// 						Result: to.Ptr("SUCCESS"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
