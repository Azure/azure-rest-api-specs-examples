package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v4"
)

// Generated from example definition: 2026-10-01/ValidateProbe.json
func ExampleManagementClient_ValidateProbe() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ValidateProbe(ctx, armcdn.ValidateProbeInput{
		ProbeURL: to.Ptr("https://www.bing.com/image"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.ManagementClientValidateProbeResponse{
	// 	ValidateProbeOutput: armcdn.ValidateProbeOutput{
	// 		ErrorCode: to.Ptr("None"),
	// 		IsValid: to.Ptr(true),
	// 	},
	// }
}
