package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/SupportedOperatingSystems_Get.json
func ExampleSupportedOperatingSystemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSupportedOperatingSystemsClient().Get(ctx, "resourceGroupPS1", "vault1", &armrecoveryservicessiterecovery.SupportedOperatingSystemsClientGetOptions{InstanceType: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SupportedOperatingSystems = armrecoveryservicessiterecovery.SupportedOperatingSystems{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationSupportedOperatingSystems"),
	// 	ID: to.Ptr("/Subscriptions/bc403605-c2b0-43dd-abe9-0162124b1ee1/resourceGroups/oneBoxRG/providers/Microsoft.RecoveryServices/vaults/oneBoxRSVault/replicationSupportedOperatingSystems/Default"),
	// 	Properties: &armrecoveryservicessiterecovery.SupportedOSProperties{
	// 		SupportedOsList: []*armrecoveryservicessiterecovery.SupportedOSProperty{
	// 			{
	// 				InstanceType: to.Ptr("A2A"),
	// 				SupportedOs: []*armrecoveryservicessiterecovery.SupportedOSDetails{
	// 					{
	// 						OSName: to.Ptr("centos"),
	// 						OSType: to.Ptr("linux"),
	// 						OSVersions: []*armrecoveryservicessiterecovery.OSVersionWrapper{
	// 							{
	// 								Version: to.Ptr("6.0"),
	// 						}},
	// 					},
	// 					{
	// 						OSName: to.Ptr("Windows Server 2008 R2 Datacenter"),
	// 						OSType: to.Ptr("windows"),
	// 						OSVersions: []*armrecoveryservicessiterecovery.OSVersionWrapper{
	// 							{
	// 								ServicePack: to.Ptr("1"),
	// 								Version: to.Ptr("6.1"),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// }
}
