package armanalysisservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/listSkusForNew.json
func ExampleServersClient_ListSKUsForNew() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armanalysisservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().ListSKUsForNew(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUEnumerationForNewResourceResult = armanalysisservices.SKUEnumerationForNewResourceResult{
	// 	Value: []*armanalysisservices.ResourceSKU{
	// 		{
	// 			Name: to.Ptr("B1"),
	// 		},
	// 		{
	// 			Name: to.Ptr("B2"),
	// 		},
	// 		{
	// 			Name: to.Ptr("D1"),
	// 		},
	// 		{
	// 			Name: to.Ptr("S0"),
	// 		},
	// 		{
	// 			Name: to.Ptr("S1"),
	// 		},
	// 		{
	// 			Name: to.Ptr("S2"),
	// 		},
	// 		{
	// 			Name: to.Ptr("S3"),
	// 		},
	// 		{
	// 			Name: to.Ptr("S4"),
	// 	}},
	// }
}
