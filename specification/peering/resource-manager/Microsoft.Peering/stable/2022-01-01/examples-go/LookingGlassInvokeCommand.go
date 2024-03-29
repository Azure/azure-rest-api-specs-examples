package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/LookingGlassInvokeCommand.json
func ExampleLookingGlassClient_Invoke() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLookingGlassClient().Invoke(ctx, armpeering.LookingGlassCommandTraceroute, armpeering.LookingGlassSourceTypeAzureRegion, "West US", "0.0.0.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LookingGlassOutput = armpeering.LookingGlassOutput{
	// 	Command: to.Ptr(armpeering.CommandTraceroute),
	// 	Output: to.Ptr("traceroute to 0.0.0.0, 64 hops max, 52 bytes packets\n 1 West US (1.1.1.1) 0.111ms 0.222ms 0.333ms"),
	// }
}
