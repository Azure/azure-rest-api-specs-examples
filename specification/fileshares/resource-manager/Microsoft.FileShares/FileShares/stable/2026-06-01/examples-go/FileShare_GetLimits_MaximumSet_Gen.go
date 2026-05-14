package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/FileShare_GetLimits_MaximumSet_Gen.json
func ExampleInformationalOperationsClient_GetLimits_fileShareGetLimitsMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInformationalOperationsClient().GetLimits(ctx, "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armfileshares.InformationalOperationsClientGetLimitsResponse{
	// 	FileShareLimitsResponse: armfileshares.FileShareLimitsResponse{
	// 		Properties: &armfileshares.FileShareLimitsOutput{
	// 			Limits: &armfileshares.FileShareLimits{
	// 				MaxFileShares: to.Ptr[int32](21),
	// 				MaxFileShareSnapshots: to.Ptr[int32](5),
	// 				MaxFileShareSubnets: to.Ptr[int32](11),
	// 				MaxFileSharePrivateEndpointConnections: to.Ptr[int32](17),
	// 				MinProvisionedStorageGiB: to.Ptr[int32](15),
	// 				MaxProvisionedStorageGiB: to.Ptr[int32](18),
	// 				MinProvisionedIOPerSec: to.Ptr[int32](21),
	// 				MaxProvisionedIOPerSec: to.Ptr[int32](2),
	// 				MinProvisionedThroughputMiBPerSec: to.Ptr[int32](30),
	// 				MaxProvisionedThroughputMiBPerSec: to.Ptr[int32](14),
	// 			},
	// 			ProvisioningConstants: &armfileshares.FileShareProvisioningConstants{
	// 				BaseIOPerSec: to.Ptr[int32](9),
	// 				ScalarIOPerSec: to.Ptr[float64](17),
	// 				BaseThroughputMiBPerSec: to.Ptr[int32](17),
	// 				ScalarThroughputMiBPerSec: to.Ptr[float64](12),
	// 				GuardrailIOPerSecScalar: to.Ptr[float64](5),
	// 				GuardrailThroughputScalar: to.Ptr[float64](5),
	// 			},
	// 		},
	// 	},
	// }
}
