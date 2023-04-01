package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateVMMServer.json
func ExampleVmmServersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVmmServersClient().BeginCreateOrUpdate(ctx, "testrg", "ContosoVMMServer", armscvmm.VMMServer{
		ExtendedLocation: &armscvmm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
			Type: to.Ptr("customLocation"),
		},
		Location: to.Ptr("East US"),
		Properties: &armscvmm.VMMServerProperties{
			Credentials: &armscvmm.VMMServerPropertiesCredentials{
				Password: to.Ptr("password"),
				Username: to.Ptr("testuser"),
			},
			Fqdn: to.Ptr("VMM.contoso.com"),
			Port: to.Ptr[int32](1234),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMMServer = armscvmm.VMMServer{
	// 	Name: to.Ptr("ContosoVMMServer"),
	// 	Type: to.Ptr("Microsoft.SCVMM/VMMServers"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armscvmm.VMMServerProperties{
	// 		ConnectionStatus: to.Ptr("Connected"),
	// 		Fqdn: to.Ptr("VMM.contoso.com"),
	// 		Port: to.Ptr[int32](1234),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UUID: to.Ptr("fd3c3665-1729-4b7b-9a38-238e83b0f98b"),
	// 		Version: to.Ptr("2.0"),
	// 	},
	// }
}
