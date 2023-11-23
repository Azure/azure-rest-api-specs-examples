package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/licenseProfile/LicenseProfile_Get.json
func ExampleLicenseProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLicenseProfilesClient().Get(ctx, "myResourceGroup", "myMachine", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 			EsuKeys: []*armhybridcompute.EsuKey{
	// 				{
	// 					LicenseStatus: to.Ptr("licenseStatus1"),
	// 					SKU: to.Ptr("skuNumber1"),
	// 				},
	// 				{
	// 					LicenseStatus: to.Ptr("licenseStatus2"),
	// 					SKU: to.Ptr("skuNumber2"),
	// 			}},
	// 			EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityEligible),
	// 			EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
	// 			ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
	// 			AssignedLicense: to.Ptr("{LicenseResourceId}"),
	// 		},
	// 	},
	// }
}
