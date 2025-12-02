package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PaloAltoNetworksCloudngfwOperations_listCloudManagerTenants_MinimumSet_Gen.json
func ExamplePaloAltoNetworksCloudngfwOperationsClient_ListCloudManagerTenants_paloAltoNetworksCloudngfwOperationsListCloudManagerTenantsMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaloAltoNetworksCloudngfwOperationsClient().ListCloudManagerTenants(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.PaloAltoNetworksCloudngfwOperationsClientListCloudManagerTenantsResponse{
	// 	CloudManagerTenantList: &armpanngfw.CloudManagerTenantList{
	// 		Value: []*string{
	// 			to.Ptr("TenantName - 1234"),
	// 		},
	// 	},
	// }
}
