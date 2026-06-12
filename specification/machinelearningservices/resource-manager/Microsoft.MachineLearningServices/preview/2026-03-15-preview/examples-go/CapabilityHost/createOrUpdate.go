package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/CapabilityHost/createOrUpdate.json
func ExampleCapabilityHostsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapabilityHostsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "capabilityHostName", armmachinelearning.CapabilityHost{
		Properties: &armmachinelearning.CapabilityHostProperties{
			AcaEnvironmentConnections: []*string{
				to.Ptr("sampleAcaEnvironmentConnection"),
			},
			AiServicesConnections: []*string{
				to.Ptr("sampleAIServiceConnection"),
			},
			CustomerSubnet: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroups/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
			StorageConnections: []*string{
				to.Ptr("sampleStorageConnection"),
			},
			ThreadStorageConnections: []*string{
				to.Ptr("sampleThreadStorageConnection"),
			},
			VectorStoreConnections: []*string{
				to.Ptr("sampleVectorStoreConnection"),
			},
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
	// res = armmachinelearning.CapabilityHostsClientCreateOrUpdateResponse{
	// 	CapabilityHost: armmachinelearning.CapabilityHost{
	// 		Name: to.Ptr("capabilityHostName"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/capabilityHosts"),
	// 		ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/capabilityHosts/capabilityHostName"),
	// 		Properties: &armmachinelearning.CapabilityHostProperties{
	// 			Description: to.Ptr("string"),
	// 			AcaEnvironmentConnections: []*string{
	// 				to.Ptr("sampleAcaEnvironmentConnection"),
	// 			},
	// 			AiServicesConnections: []*string{
	// 				to.Ptr("sampleAIServiceConnection"),
	// 			},
	// 			CustomerSubnet: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroups/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubne"),
	// 			ProvisioningState: to.Ptr(armmachinelearning.CapabilityHostProvisioningStateSucceeded),
	// 			StorageConnections: []*string{
	// 				to.Ptr("sampleStorageConnection"),
	// 			},
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			ThreadStorageConnections: []*string{
	// 				to.Ptr("sampleThreadStorageConnection"),
	// 			},
	// 			VectorStoreConnections: []*string{
	// 				to.Ptr("sampleVectoStoreConnection"),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:22"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:22"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 		},
	// 	},
	// }
}
