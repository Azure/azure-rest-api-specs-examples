package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/volumes/create_update.json
func ExampleVolumeClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmesh.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeClient().Create(ctx, "sbz_demo", "sampleVolume", armservicefabricmesh.VolumeResourceDescription{
		Location: to.Ptr("EastUS"),
		Tags:     map[string]*string{},
		Properties: &armservicefabricmesh.VolumeResourceProperties{
			Description: to.Ptr("Service Fabric Mesh sample volume."),
			AzureFileParameters: &armservicefabricmesh.VolumeProviderParametersAzureFile{
				AccountKey:  to.Ptr("provide-account-key-here"),
				AccountName: to.Ptr("sbzdemoaccount"),
				ShareName:   to.Ptr("sharel"),
			},
			Provider: to.Ptr(armservicefabricmesh.VolumeProviderSFAzureFile),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeResourceDescription = armservicefabricmesh.VolumeResourceDescription{
	// 	Name: to.Ptr("sampleVolume"),
	// 	Type: to.Ptr("Microsoft.ServiceFabricMesh/volumes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/volumes/sampleVolume"),
	// 	Location: to.Ptr("EastUS"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabricmesh.VolumeResourceProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Description: to.Ptr("Service Fabric Mesh sample volume."),
	// 		AzureFileParameters: &armservicefabricmesh.VolumeProviderParametersAzureFile{
	// 			AccountName: to.Ptr("sbzdemoaccount"),
	// 			ShareName: to.Ptr("sharel"),
	// 		},
	// 		Provider: to.Ptr(armservicefabricmesh.VolumeProviderSFAzureFile),
	// 		Status: to.Ptr(armservicefabricmesh.ResourceStatusReady),
	// 	},
	// }
}
