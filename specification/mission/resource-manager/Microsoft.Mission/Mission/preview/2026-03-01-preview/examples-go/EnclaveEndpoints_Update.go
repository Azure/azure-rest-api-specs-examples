package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/EnclaveEndpoints_Update.json
func ExampleEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("73CEECEF-2C30-488E-946F-D20F414D99BA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointsClient().BeginUpdate(ctx, "rgopenapi", "TestMyEnclave", "TestMyEnclaveEndpoint", armenclave.EndpointPatchModel{
		Tags: map[string]*string{
			"sampletag": to.Ptr("samplevalue"),
		},
		Properties: &armenclave.EndpointPatchProperties{
			RuleCollection: []*armenclave.EndpointDestinationRule{
				{
					EndpointRuleName: to.Ptr("54CEECEF-2C30-488E-946F-D20F414D99BA"),
					Destination:      to.Ptr("10.0.0.0/24"),
					Ports:            to.Ptr("443"),
					Protocols: []*armenclave.EndpointProtocol{
						to.Ptr(armenclave.EndpointProtocolTCP),
					},
				},
			},
			UpdateMode: to.Ptr(armenclave.UpdateModeAutomatic),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.EndpointsClientUpdateResponse{
	// 	EndpointResource: armenclave.EndpointResource{
	// 		Properties: &armenclave.EndpointProperties{
	// 			RuleCollection: []*armenclave.EndpointDestinationRule{
	// 				{
	// 					EndpointRuleName: to.Ptr("54CEECEF-2C30-488E-946F-D20F414D99BA"),
	// 					Destination: to.Ptr("10.0.0.0/24"),
	// 					Ports: to.Ptr("443"),
	// 					Protocols: []*armenclave.EndpointProtocol{
	// 						to.Ptr(armenclave.EndpointProtocolTCP),
	// 					},
	// 				},
	// 			},
	// 			UpdateMode: to.Ptr(armenclave.UpdateModeAutomatic),
	// 			ProvisioningState: to.Ptr(armenclave.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"sampletag": to.Ptr("samplevalue"),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/virtualenclaves/TestMyEnclave/enclaveendpoints/TestMyEnclaveEndpoint"),
	// 		Name: to.Ptr("TestMyEnclaveEndpoint"),
	// 		Type: to.Ptr("Microsoft.Mission/virtualenclaves/enclaveendpoints"),
	// 		SystemData: &armenclave.SystemData{
	// 			CreatedBy: to.Ptr("myAlias"),
	// 			CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("myAlias"),
	// 			LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 		},
	// 	},
	// }
}
