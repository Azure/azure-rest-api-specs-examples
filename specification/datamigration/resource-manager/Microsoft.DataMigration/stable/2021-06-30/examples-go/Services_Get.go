package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Services_Get.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "DmsSdkRg", "DmsSdkService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armdatamigration.Service{
	// 	Name: to.Ptr("DmsSdkService"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Etag: to.Ptr("qt85+bWDN84/6PR8Gllxf63krZcXZX1h3wxAbs6pCjc="),
	// 	Properties: &armdatamigration.ServiceProperties{
	// 		ProvisioningState: to.Ptr(armdatamigration.ServiceProvisioningStateSucceeded),
	// 		PublicKey: to.Ptr("0;AQAB;sCMXX00QDnuOIEQdXyY96h3VPSk3ccsBeYOGxntqg7y6oXbq4JbQHf8GnUhf5NsmUYBq1rO+6xSumIuLLANIIqNPpHCT8x7W+tcuShlju4fdSKfpyKbJnTJJzGLhRJYG97hDPo9D76qpEYQN66k0IfkKJVDTh7LAxqcE3lxdA98LB2EDM6vYoZ+raKXygYc1UUHDpkfksLk2YRTVCaayTFBh5HT/+CY+48V+rbx/nruvEEEbSg720m5Fc0k8ivM3A+RGkeK9pzQx+TtskSCDXiWL50a8Rrb5ORkBusgY+s0XLyKOuxZCIUOEO8JzUjCeZYdOfLJzx8yfTxEmSlOPqQ=="),
	// 		VirtualSubnetID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkTestNetwork/providers/Microsoft.Network/virtualNetworks/DmsSdkTestNetwork/subnets/default"),
	// 	},
	// 	SKU: &armdatamigration.ServiceSKU{
	// 		Name: to.Ptr("Basic_1vCore"),
	// 		Size: to.Ptr("1 vCore"),
	// 		Tier: to.Ptr("Basic"),
	// 	},
	// }
}
