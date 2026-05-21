package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/Services_CreateOrUpdate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "DmsSdkRg", "DmsSdkService", armdatamigration.Service{
		Location: to.Ptr("southcentralus"),
		Properties: &armdatamigration.ServiceProperties{
			VirtualSubnetID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkTestNetwork/providers/Microsoft.Network/virtualNetworks/DmsSdkTestNetwork/subnets/default"),
		},
		SKU: &armdatamigration.ServiceSKU{
			Name: to.Ptr("Basic_1vCore"),
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
	// res = armdatamigration.ServicesClientCreateOrUpdateResponse{
	// 	Service: armdatamigration.Service{
	// 		Name: to.Ptr("DmsSdkService"),
	// 		Type: to.Ptr("Microsoft.DataMigration/services"),
	// 		Etag: to.Ptr("9QuK/U0GvTJpjIrlOzITXVy68+bmmQ3bFkHudLxmkUw="),
	// 		ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService"),
	// 		Location: to.Ptr("southcentralus"),
	// 		Properties: &armdatamigration.ServiceProperties{
	// 			ProvisioningState: to.Ptr(armdatamigration.ServiceProvisioningStateAccepted),
	// 			VirtualSubnetID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkTestNetwork/providers/Microsoft.Network/virtualNetworks/DmsSdkTestNetwork/subnets/default"),
	// 		},
	// 		SKU: &armdatamigration.ServiceSKU{
	// 			Name: to.Ptr("Basic_1vCore"),
	// 			Size: to.Ptr("1 vCore"),
	// 			Tier: to.Ptr("Basic"),
	// 		},
	// 	},
	// }
}
