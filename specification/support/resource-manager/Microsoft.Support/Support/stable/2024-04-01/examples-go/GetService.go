package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: 2024-04-01/GetService.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "service_guid", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsupport.ServicesClientGetResponse{
	// 	Service: armsupport.Service{
	// 		Name: to.Ptr("service_guid"),
	// 		Type: to.Ptr("Microsoft.Support/services"),
	// 		ID: to.Ptr("/providers/Microsoft.Support/services/service_guid"),
	// 		Properties: &armsupport.ServiceProperties{
	// 			DisplayName: to.Ptr("Virtual Machine running Windows"),
	// 			ResourceTypes: []*string{
	// 				to.Ptr("MICROSOFT.CLASSICCOMPUTE/VIRTUALMACHINES"),
	// 				to.Ptr("MICROSOFT.COMPUTE/VIRTUALMACHINES"),
	// 			},
	// 		},
	// 	},
	// }
}
