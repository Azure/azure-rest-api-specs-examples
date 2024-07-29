package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a651ba25cda4eec698a3a4e35f867ecc2681d126/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ManagedClusterVersionGet_example.json
func ExampleManagedClusterVersionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClusterVersionClient().Get(ctx, "eastus", "7.2.477.9590", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedClusterCodeVersionResult = armservicefabricmanagedclusters.ManagedClusterCodeVersionResult{
	// 	Name: to.Ptr("7.2.477.9590"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/locations/environments/managedClusterVersions"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/environments/Windows/managedClusterVersions/7.2.477.9590"),
	// 	Properties: &armservicefabricmanagedclusters.ManagedClusterVersionDetails{
	// 		ClusterCodeVersion: to.Ptr("7.2.477.9590"),
	// 		OSType: to.Ptr(armservicefabricmanagedclusters.OsTypeWindows),
	// 		SupportExpiryUTC: to.Ptr("2021-11-30T00:00:00"),
	// 	},
	// }
}
