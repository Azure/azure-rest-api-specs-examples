package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListByEdgeZone_virtualMachineImagesEdgeZoneListByEdgeZoneMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListByEdgeZone(ctx, "WestUS", "microsoftlosangeles1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMImagesInEdgeZoneListResult = armcompute.VMImagesInEdgeZoneListResult{
	// 	Value: []*armcompute.VirtualMachineImageResource{
	// 		{
	// 			ID: to.Ptr("/Subscriptions/5ece5940-d962-4dad-a98f-ca9ac0f021a5/Providers/Microsoft.Compute/Locations/westus/Publishers/CANONICAL/ArtifactTypes/VMImage/Offers/UBUNTUSERVER/Skus/18_04-LTS-GEN2/Versions/18.04.202107200"),
	// 			Name: to.Ptr("18.04.202107200"),
	// 			ExtendedLocation: &armcompute.ExtendedLocation{
	// 				Name: to.Ptr("microsoftlosangeles1"),
	// 				Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 			},
	// 			Location: to.Ptr("WestUS"),
	// 	}},
	// }
}
