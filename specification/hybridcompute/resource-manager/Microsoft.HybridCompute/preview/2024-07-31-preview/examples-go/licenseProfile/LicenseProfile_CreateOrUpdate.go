package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/licenseProfile/LicenseProfile_CreateOrUpdate.json
func ExampleLicenseProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLicenseProfilesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMachine", armhybridcompute.LicenseProfile{
		Location: to.Ptr("eastus2euap"),
		Properties: &armhybridcompute.LicenseProfileProperties{
			EsuProfile: &armhybridcompute.LicenseProfileArmEsuProperties{
				AssignedLicense: to.Ptr("{LicenseResourceId}"),
			},
			ProductProfile: &armhybridcompute.LicenseProfileArmProductProfileProperties{
				ProductFeatures: []*armhybridcompute.ProductFeature{
					{
						Name:               to.Ptr("Hotpatch"),
						SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
					}},
				ProductType:        to.Ptr(armhybridcompute.LicenseProfileProductTypeWindowsServer),
				SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
			},
			SoftwareAssurance: &armhybridcompute.LicenseProfilePropertiesSoftwareAssurance{
				SoftwareAssuranceCustomer: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LicenseProfile = armhybridcompute.LicenseProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines/licenseProfiles"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/licenseProfiles/default"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.LicenseProfileProperties{
	// 		EsuProfile: &armhybridcompute.LicenseProfileArmEsuProperties{
	// 			AssignedLicenseImmutableID: to.Ptr("{Guid}"),
	// 			EsuKeys: []*armhybridcompute.EsuKey{
	// 				{
	// 					LicenseStatus: to.Ptr[int32](1),
	// 					SKU: to.Ptr("skuNumber1"),
	// 				},
	// 				{
	// 					LicenseStatus: to.Ptr[int32](1),
	// 					SKU: to.Ptr("skuNumber2"),
	// 			}},
	// 			EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityEligible),
	// 			EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
	// 			ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
	// 			AssignedLicense: to.Ptr("{LicenseResourceId}"),
	// 		},
	// 		ProductProfile: &armhybridcompute.LicenseProfileArmProductProfileProperties{
	// 			BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 			BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 			DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 			EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 			ProductFeatures: []*armhybridcompute.ProductFeature{
	// 				{
	// 					Name: to.Ptr("Hotpatch"),
	// 					BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 					BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 					DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 					EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 					SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabling),
	// 			}},
	// 			ProductType: to.Ptr(armhybridcompute.LicenseProfileProductTypeWindowsServer),
	// 			SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabling),
	// 		},
	// 		SoftwareAssurance: &armhybridcompute.LicenseProfilePropertiesSoftwareAssurance{
	// 			SoftwareAssuranceCustomer: to.Ptr(true),
	// 		},
	// 	},
	// }
}
