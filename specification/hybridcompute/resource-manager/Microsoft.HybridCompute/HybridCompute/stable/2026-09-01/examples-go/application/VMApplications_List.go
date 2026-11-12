package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v3"
)

// Generated from example definition: 2026-09-01/application/VMApplications_List.json
func ExampleVMApplicationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVMApplicationsClient().NewListPager("myResourceGroup", "myMachine", &armhybridcompute.VMApplicationsClientListOptions{
		Top: to.Ptr[int32](100)})
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
		// page = armhybridcompute.VMApplicationsClientListResponse{
		// 	VMApplicationsListResult: armhybridcompute.VMApplicationsListResult{
		// 		Value: []*armhybridcompute.VMApplication{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine/applications/myApplication"),
		// 				Name: to.Ptr("myApplication"),
		// 				Type: to.Ptr("Microsoft.HybridCompute/machines/applications"),
		// 				Location: to.Ptr("eastus2euap"),
		// 				Properties: &armhybridcompute.VMApplicationProperties{
		// 					Order: to.Ptr[int32](1),
		// 					PackageReferenceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/galleryResourceGroup/providers/Microsoft.Compute/galleries/myGallery/applications/myApplication/versions/1.0.0"),
		// 					TreatFailureAsDeploymentFailure: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine/applications?api-version=2026-09-01&$skipToken=opaque-token"),
		// 	},
		// }
	}
}
