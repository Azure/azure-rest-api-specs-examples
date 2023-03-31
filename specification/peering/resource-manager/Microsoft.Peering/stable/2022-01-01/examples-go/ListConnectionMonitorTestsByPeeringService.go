package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListConnectionMonitorTestsByPeeringService.json
func ExampleConnectionMonitorTestsClient_NewListByPeeringServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionMonitorTestsClient().NewListByPeeringServicePager("rgName", "peeringServiceName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ConnectionMonitorTestListResult = armpeering.ConnectionMonitorTestListResult{
		// 	Value: []*armpeering.ConnectionMonitorTest{
		// 		{
		// 			Name: to.Ptr("connectionMonitorTestName1"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName/connectionMonitorTests/connectionMonitorTestName1"),
		// 			Properties: &armpeering.ConnectionMonitorTestProperties{
		// 				Path: []*string{
		// 					to.Ptr("source"),
		// 					to.Ptr("hop1"),
		// 					to.Ptr("hop2"),
		// 					to.Ptr("destination")},
		// 					Destination: to.Ptr("Example Destination"),
		// 					DestinationPort: to.Ptr[int32](443),
		// 					IsTestSuccessful: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 					SourceAgent: to.Ptr("Example Source Agent"),
		// 					TestFrequencyInSec: to.Ptr[int32](30),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("connectionMonitorTestName2"),
		// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName/connectionMonitorTests/connectionMonitorTestName2"),
		// 				Properties: &armpeering.ConnectionMonitorTestProperties{
		// 					Path: []*string{
		// 						to.Ptr("source"),
		// 						to.Ptr("hop1"),
		// 						to.Ptr("hop2"),
		// 						to.Ptr("destination")},
		// 						Destination: to.Ptr("Example Destination 2"),
		// 						DestinationPort: to.Ptr[int32](443),
		// 						IsTestSuccessful: to.Ptr(false),
		// 						ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 						SourceAgent: to.Ptr("Example Source Agent"),
		// 						TestFrequencyInSec: to.Ptr[int32](30),
		// 					},
		// 			}},
		// 		}
	}
}
