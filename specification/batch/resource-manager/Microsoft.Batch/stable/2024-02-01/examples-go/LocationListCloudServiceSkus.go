package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/LocationListCloudServiceSkus.json
func ExampleLocationClient_NewListSupportedCloudServiceSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListSupportedCloudServiceSKUsPager("japaneast", &armbatch.LocationClientListSupportedCloudServiceSKUsOptions{Maxresults: nil,
		Filter: nil,
	})
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
		// page.SupportedSKUsResult = armbatch.SupportedSKUsResult{
		// 	Value: []*armbatch.SupportedSKU{
		// 		{
		// 			Name: to.Ptr("Small"),
		// 			Capabilities: []*armbatch.SKUCapability{
		// 				{
		// 					Name: to.Ptr("MaxResourceVolumeMB"),
		// 					Value: to.Ptr("20480"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUs"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("HyperVGenerations"),
		// 					Value: to.Ptr("V1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryGB"),
		// 					Value: to.Ptr("0.75"),
		// 				},
		// 				{
		// 					Name: to.Ptr("LowPriorityCapable"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUsAvailable"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("EphemeralOSDiskSupported"),
		// 					Value: to.Ptr("False"),
		// 			}},
		// 			FamilyName: to.Ptr("standardA0_A7Family"),
		// 	}},
		// }
	}
}
