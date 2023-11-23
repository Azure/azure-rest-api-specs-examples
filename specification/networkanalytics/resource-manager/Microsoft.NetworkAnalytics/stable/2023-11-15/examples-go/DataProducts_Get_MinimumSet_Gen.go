package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_Get_MinimumSet_Gen.json
func ExampleDataProductsClient_Get_dataProductsGetMaximumSetGenGeneratedByMinimumSetRuleMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataProductsClient().Get(ctx, "aoiresourceGroupName", "dataproduct01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataProduct = armnetworkanalytics.DataProduct{
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/aoiresourceGroupName/providers/Microsoft.NetworkAnalytics/DataProducts/dataproduct01"),
	// 	Location: to.Ptr("eastus"),
	// }
}
