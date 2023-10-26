package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Endpoints_ListResourceUsage.json
func ExampleEndpointsClient_NewListResourceUsagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEndpointsClient().NewListResourceUsagePager("RG", "profile1", "endpoint1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ResourceUsageListResult = armcdn.ResourceUsageListResult{
		// 	Value: []*armcdn.ResourceUsage{
		// 		{
		// 			CurrentValue: to.Ptr[int32](1),
		// 			Limit: to.Ptr[int32](20),
		// 			ResourceType: to.Ptr("customdomain"),
		// 			Unit: to.Ptr(armcdn.ResourceUsageUnitCount),
		// 		},
		// 		{
		// 			CurrentValue: to.Ptr[int32](0),
		// 			Limit: to.Ptr[int32](25),
		// 			ResourceType: to.Ptr("geofilter"),
		// 			Unit: to.Ptr(armcdn.ResourceUsageUnitCount),
		// 	}},
		// }
	}
}
