package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_InvokeAvailabilityZoneDetails_northeurope.json
func ExampleSAPVirtualInstancesClient_GetAvailabilityZoneDetails_sapAvailabilityZoneDetailsInNorthEurope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPVirtualInstancesClient().GetAvailabilityZoneDetails(ctx, "northeurope", armworkloadssapvirtualinstance.SAPAvailabilityZoneDetailsRequest{
		AppLocation:  to.Ptr("northeurope"),
		SapProduct:   to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		DatabaseType: to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientGetAvailabilityZoneDetailsResponse{
	// 	SAPAvailabilityZoneDetailsResult: &armworkloadssapvirtualinstance.SAPAvailabilityZoneDetailsResult{
	// 		AvailabilityZonePairs: []*armworkloadssapvirtualinstance.SAPAvailabilityZonePair{
	// 			{
	// 				ZoneA: to.Ptr[int64](2),
	// 				ZoneB: to.Ptr[int64](3),
	// 			},
	// 		},
	// 	},
	// }
}
