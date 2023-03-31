package armoep_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2022-04-04-preview/examples/EnergyServices_ListPartitions_MaximumSet_Gen.json
func ExampleEnergyServicesClient_ListPartitions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoep.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnergyServicesClient().ListPartitions(ctx, "rgoep", "aaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataPartitionsListResult = armoep.DataPartitionsListResult{
	// 	Value: []*armoep.DataPartitionProperties{
	// 		{
	// 			Name: to.Ptr("aaaaaaaaaaa"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 	}},
	// }
}
