package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v3"
)

// Generated from example definition: 2026-09-01/licenseProfile/LicenseProfile_List.json
func ExampleLicenseProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLicenseProfilesClient().NewListPager("myResourceGroup", "myMachine", nil)
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
		// page = armhybridcompute.LicenseProfilesClientListResponse{
		// 	LicenseProfilesListResult: armhybridcompute.LicenseProfilesListResult{
		// 		Value: []*armhybridcompute.LicenseProfile{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.HybridCompute/machines/licenseProfiles"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/licenseProfiles/default"),
		// 				Location: to.Ptr("eastus2euap"),
		// 				Properties: &armhybridcompute.LicenseProfileProperties{
		// 					EsuProfile: &armhybridcompute.LicenseProfileArmEsuProperties{
		// 						AssignedLicense: to.Ptr("{LicenseResourceId}"),
		// 						EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityEligible),
		// 						EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
		// 						EsuKeys: []*armhybridcompute.EsuKey{
		// 							{
		// 								LicenseStatus: to.Ptr[int32](1),
		// 								SKU: to.Ptr("skuNumber1"),
		// 							},
		// 							{
		// 								LicenseStatus: to.Ptr[int32](1),
		// 								SKU: to.Ptr("skuNumber2"),
		// 							},
		// 						},
		// 						ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
		// 					},
		// 					ProductProfile: &armhybridcompute.LicenseProfileArmProductProfileProperties{
		// 						BillingEndDate: to.Ptr(time.Date(2023, time.November, 5, 20, 36, 49, 745000000, time.UTC)),
		// 						BillingStartDate: to.Ptr(time.Date(2023, time.October, 5, 20, 36, 49, 745000000, time.UTC)),
		// 						DisenrollmentDate: to.Ptr(time.Date(2023, time.November, 5, 20, 36, 49, 745000000, time.UTC)),
		// 						EnrollmentDate: to.Ptr(time.Date(2023, time.October, 5, 20, 36, 49, 745000000, time.UTC)),
		// 						ProductFeatures: []*armhybridcompute.ProductFeature{
		// 							{
		// 								Name: to.Ptr("Hotpatch"),
		// 								BillingEndDate: to.Ptr(time.Date(2023, time.November, 5, 20, 36, 49, 745000000, time.UTC)),
		// 								BillingStartDate: to.Ptr(time.Date(2023, time.October, 5, 20, 36, 49, 745000000, time.UTC)),
		// 								DisenrollmentDate: to.Ptr(time.Date(2023, time.November, 5, 20, 36, 49, 745000000, time.UTC)),
		// 								EnrollmentDate: to.Ptr(time.Date(2023, time.October, 5, 20, 36, 49, 745000000, time.UTC)),
		// 								SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
		// 							},
		// 						},
		// 						ProductType: to.Ptr(armhybridcompute.LicenseProfileProductTypeWindowsServer),
		// 						SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
		// 					},
		// 					SoftwareAssurance: &armhybridcompute.LicenseProfilePropertiesSoftwareAssurance{
		// 						SoftwareAssuranceCustomer: to.Ptr(true),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
