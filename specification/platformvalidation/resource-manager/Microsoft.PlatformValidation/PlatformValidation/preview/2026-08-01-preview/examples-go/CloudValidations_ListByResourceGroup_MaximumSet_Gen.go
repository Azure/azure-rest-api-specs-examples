package armplatformvalidation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/platformvalidation/armplatformvalidation"
)

// Generated from example definition: 2026-08-01-preview/CloudValidations_ListByResourceGroup_MaximumSet_Gen.json
func ExampleCloudValidationsClient_NewListByResourceGroupPager_cloudValidationsListByResourceGroupMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplatformvalidation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudValidationsClient().NewListByResourceGroupPager("rgvalidate", nil)
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
		// 				Properties: &armplatformvalidation.CloudValidationProperties{
		// 					Description: to.Ptr("Cloud validation that groups platform validation execution plans for the target subscription."),
		// 					ProvisioningState: to.Ptr(armplatformvalidation.ProvisioningStateSucceeded),
		// 					ManagedOnBehalfOfConfiguration: &armplatformvalidation.ManagedOnBehalfOfConfiguration{
		// 						MoboBrokerResources: []*armplatformvalidation.MoboBrokerResource{
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgvalidate/providers/Microsoft.PlatformValidation/cloudValidations/cvtest01"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				Location: to.Ptr("southcentralus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgvalidate/providers/Microsoft.PlatformValidation/cloudValidations/cvtest01"),
		// 				Name: to.Ptr("cvtest01"),
		// 				Type: to.Ptr("Microsoft.PlatformValidation/cloudValidations"),
		// 				SystemData: &armplatformvalidation.SystemData{
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armplatformvalidation.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.June, 1, 11, 52, 22, 926000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armplatformvalidation.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.June, 1, 11, 52, 22, 926000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PlatformValidation/validationTests?api-version=2026-08-01-preview&$skiptoken=cD0yMDI2LTA3LTAx"),
		// 	},
		// }
	}
}
