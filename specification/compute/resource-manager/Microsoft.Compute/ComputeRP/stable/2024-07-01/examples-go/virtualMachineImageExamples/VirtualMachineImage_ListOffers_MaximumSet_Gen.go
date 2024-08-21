package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListOffers_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListOffers_virtualMachineImageListOffersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListOffers(ctx, "aaaaaaa", "aaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}
