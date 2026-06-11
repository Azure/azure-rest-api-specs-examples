package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2026-05-01-preview/ManagedApplyMaintenanceWindowPost_example.json
func ExampleManagedApplyMaintenanceWindowClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedApplyMaintenanceWindowClient().Post(ctx, "resourceGroup1", "mycluster1", &armservicefabricmanagedclusters.ManagedApplyMaintenanceWindowClientPostOptions{
		Body: &armservicefabricmanagedclusters.ApplyMaintenanceWindowRequest{
			StartDateTime: to.Ptr("2026-04-07 13:00"),
			Duration:      to.Ptr("08:30"),
			TimeZone:      to.Ptr("Pacific Standard Time"),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.ManagedApplyMaintenanceWindowClientPostResponse{
	// }
}
