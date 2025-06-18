package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/licenseProfile/LicenseProfile_List.json
func ExampleLicenseProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.LicenseProfilesListResult = armhybridcompute.LicenseProfilesListResult{
		// 	Value: []*armhybridcompute.LicenseProfile{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines/licenseProfiles"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/licenseProfiles/default"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armhybridcompute.LicenseProfileProperties{
		// 				EsuProfile: &armhybridcompute.LicenseProfileArmEsuProperties{
		// 					EsuKeys: []*armhybridcompute.EsuKey{
		// 						{
		// 							LicenseStatus: to.Ptr[int32](1),
		// 							SKU: to.Ptr("skuNumber1"),
		// 						},
		// 						{
		// 							LicenseStatus: to.Ptr[int32](1),
		// 							SKU: to.Ptr("skuNumber2"),
		// 					}},
		// 					EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityEligible),
		// 					EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
		// 					ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
		// 					AssignedLicense: to.Ptr("{LicenseResourceId}"),
		// 				},
		// 				ProductProfile: &armhybridcompute.LicenseProfileArmProductProfileProperties{
		// 					BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
		// 					BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
		// 					DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
		// 					EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
		// 					ProductFeatures: []*armhybridcompute.ProductFeature{
		// 						{
		// 							Name: to.Ptr("Hotpatch"),
		// 							BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
		// 							BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
		// 							DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
		// 							EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
		// 							SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
		// 					}},
		// 					ProductType: to.Ptr(armhybridcompute.LicenseProfileProductTypeWindowsServer),
		// 					SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
		// 				},
		// 				SoftwareAssurance: &armhybridcompute.LicenseProfilePropertiesSoftwareAssurance{
		// 					SoftwareAssuranceCustomer: to.Ptr(true),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
