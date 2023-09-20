package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateInstancePoolMin.json
func ExampleInstancePoolsClient_BeginCreateOrUpdate_createAnInstancePoolWithMinProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancePoolsClient().BeginCreateOrUpdate(ctx, "group1", "testIP", armsql.InstancePool{
		Location: to.Ptr("japaneast"),
		Properties: &armsql.InstancePoolProperties{
			LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
			SubnetID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
			VCores:      to.Ptr[int32](8),
		},
		SKU: &armsql.SKU{
			Name:   to.Ptr("GP_Gen5"),
			Family: to.Ptr("Gen5"),
			Tier:   to.Ptr("GeneralPurpose"),
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
	// res.InstancePool = armsql.InstancePool{
	// 	Name: to.Ptr("testIP"),
	// 	Type: to.Ptr("Microsoft.Sql/instancePools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsql.InstancePoolProperties{
	// 		LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
