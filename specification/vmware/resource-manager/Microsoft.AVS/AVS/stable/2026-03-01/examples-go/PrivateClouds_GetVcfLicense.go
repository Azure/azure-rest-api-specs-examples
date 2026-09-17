package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2026-03-01/PrivateClouds_GetVcfLicense.json
func ExamplePrivateCloudsClient_GetVcfLicense() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateCloudsClient().GetVcfLicense(ctx, "group1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.PrivateCloudsClientGetVcfLicenseResponse{
	// 	VcfLicenseClassification: &armavs.Vcf5License{
	// 		Kind: to.Ptr(armavs.VcfLicenseKindVcf5),
	// 		LicenseKey: to.Ptr("12345-12345-12345-12345-12345"),
	// 		EndDate: to.Ptr(time.Date(2025, time.December, 31, 23, 59, 59, 0, time.UTC)),
	// 		Cores: to.Ptr[int32](16),
	// 		BroadcomSiteID: to.Ptr("123456"),
	// 		BroadcomContractNumber: to.Ptr("123456"),
	// 		ProvisioningState: to.Ptr(armavs.LicenseProvisioningStateSucceeded),
	// 	},
	// }
}
