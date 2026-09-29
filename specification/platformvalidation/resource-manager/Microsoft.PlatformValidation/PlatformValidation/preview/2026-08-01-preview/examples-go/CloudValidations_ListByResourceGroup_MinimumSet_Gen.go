package armplatformvalidation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/platformvalidation/armplatformvalidation"
)

// Generated from example definition: 2026-08-01-preview/CloudValidations_ListByResourceGroup_MinimumSet_Gen.json
func ExampleCloudValidationsClient_NewListByResourceGroupPager_cloudValidationsListByResourceGroupMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplatformvalidation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudValidationsClient().NewListByResourceGroupPager("rgplatformvalidation", nil)
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
		// page = armplatformvalidation.CloudValidationsClientListByResourceGroupResponse{
		// 	CloudValidationListResult: armplatformvalidation.CloudValidationListResult{
		// 		Value: []*armplatformvalidation.CloudValidation{
		// 			{
		// 				Location: to.Ptr("southcentralus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgplatformvalidation/providers/Microsoft.PlatformValidation/cloudValidations/cloudvalidation1"),
		// 				Name: to.Ptr("cloudvalidation1"),
		// 				Type: to.Ptr("Microsoft.PlatformValidation/cloudValidations"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
