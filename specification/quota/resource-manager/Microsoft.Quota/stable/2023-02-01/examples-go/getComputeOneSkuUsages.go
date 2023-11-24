package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getComputeOneSkuUsages.json
func ExampleUsagesClient_Get_quotasUsagesRequestForCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUsagesClient().Get(ctx, "standardNDSFamily", "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CurrentUsagesBase = armquota.CurrentUsagesBase{
	// 	Name: to.Ptr("standardNDSFamily"),
	// 	Type: to.Ptr("Microsoft.Quota/Usages"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Usages/standardNDSFamily"),
	// 	Properties: &armquota.UsagesProperties{
	// 		Name: &armquota.ResourceName{
	// 			LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
	// 			Value: to.Ptr("standardNDSFamily"),
	// 		},
	// 		IsQuotaApplicable: to.Ptr(true),
	// 		Unit: to.Ptr("Count"),
	// 		Usages: &armquota.UsagesObject{
	// 			UsagesType: to.Ptr(armquota.UsagesTypesIndividual),
	// 			Value: to.Ptr[int32](10),
	// 		},
	// 	},
	// }
}
