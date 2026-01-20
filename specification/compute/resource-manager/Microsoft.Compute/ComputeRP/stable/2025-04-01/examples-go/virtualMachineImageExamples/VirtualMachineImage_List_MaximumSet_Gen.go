package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineImageExamples/VirtualMachineImage_List_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_List_virtualMachineImageListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().List(ctx, "aaaaaaaaaaaaaaa", "aaaaaa", "aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineImagesClientListOptions{Expand: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		Top:     to.Ptr[int32](18),
		Orderby: to.Ptr("aa"),
	})
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
