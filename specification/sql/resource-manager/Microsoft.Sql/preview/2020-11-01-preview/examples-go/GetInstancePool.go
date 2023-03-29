package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetInstancePool.json
func ExampleInstancePoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstancePoolsClient().Get(ctx, "group1", "testIP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InstancePool = armsql.InstancePool{
	// 	Name: to.Ptr("testIP"),
	// 	Type: to.Ptr("Microsoft.Sql/instancePools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"a": to.Ptr("b"),
	// 	},
	// 	Properties: &armsql.InstancePoolProperties{
	// 		LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
