package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/LocationListVirtualMachineSkus.json
func ExampleLocationClient_NewListSupportedVirtualMachineSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListSupportedVirtualMachineSKUsPager("japaneast", &armbatch.LocationClientListSupportedVirtualMachineSKUsOptions{Maxresults: nil,
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
		// 			Name: to.Ptr("Standard_D1_v2"),
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
		// 			FamilyName: to.Ptr("standardDFamily"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Standard_A1"),
		// 			BatchSupportEndOfLife: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-31T00:00:00.000Z"); return t}()),
		// 			Capabilities: []*armbatch.SKUCapability{
		// 				{
		// 					Name: to.Ptr("MaxResourceVolumeMB"),
		// 					Value: to.Ptr("71680"),
		// 				},
		// 				{
		// 					Name: to.Ptr("OSVhdSizeMB"),
		// 					Value: to.Ptr("1047552"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUs"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryPreservingMaintenanceSupported"),
		// 					Value: to.Ptr("True"),
		// 				},
		// 				{
		// 					Name: to.Ptr("HyperVGenerations"),
		// 					Value: to.Ptr("V1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryGB"),
		// 					Value: to.Ptr("1.75"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MaxDataDiskCount"),
		// 					Value: to.Ptr("2"),
		// 				},
		// 				{
		// 					Name: to.Ptr("CpuArchitectureType"),
		// 					Value: to.Ptr("x64"),
		// 				},
		// 				{
		// 					Name: to.Ptr("LowPriorityCapable"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("PremiumIO"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("VMDeploymentTypes"),
		// 					Value: to.Ptr("IaaS,PaaS"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUsAvailable"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("ACUs"),
		// 					Value: to.Ptr("100"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUsPerCore"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("EphemeralOSDiskSupported"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("EncryptionAtHostSupported"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("CapacityReservationSupported"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("AcceleratedNetworkingEnabled"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("RdmaEnabled"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MaxNetworkInterfaces"),
		// 					Value: to.Ptr("2"),
		// 			}},
		// 			FamilyName: to.Ptr("standardA0_A7Family"),
		// 	}},
		// }
	}
}
