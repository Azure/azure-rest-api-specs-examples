package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/GetConnectionMonitorTest.json
func ExampleConnectionMonitorTestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionMonitorTestsClient().Get(ctx, "rgName", "peeringServiceName", "connectionMonitorTestName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionMonitorTest = armpeering.ConnectionMonitorTest{
	// 	Name: to.Ptr("connectionMonitorTestName"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName/connectionMonitorTests/connectionMonitorTestName"),
	// 	Properties: &armpeering.ConnectionMonitorTestProperties{
	// 		Path: []*string{
	// 			to.Ptr("source"),
	// 			to.Ptr("hop1"),
	// 			to.Ptr("hop2"),
	// 			to.Ptr("destination")},
	// 			Destination: to.Ptr("Example Destination"),
	// 			DestinationPort: to.Ptr[int32](443),
	// 			IsTestSuccessful: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 			SourceAgent: to.Ptr("Example Source Agent"),
	// 			TestFrequencyInSec: to.Ptr[int32](30),
	// 		},
	// 	}
}
