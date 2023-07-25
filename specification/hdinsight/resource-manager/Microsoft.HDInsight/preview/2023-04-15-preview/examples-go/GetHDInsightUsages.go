package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetHDInsightUsages.json
func ExampleLocationsClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().ListUsages(ctx, "West US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UsagesListResult = armhdinsight.UsagesListResult{
	// 	Value: []*armhdinsight.Usage{
	// 		{
	// 			Name: &armhdinsight.LocalizedName{
	// 				LocalizedValue: to.Ptr("Cores"),
	// 				Value: to.Ptr("cores"),
	// 			},
	// 			CurrentValue: to.Ptr[int64](0),
	// 			Limit: to.Ptr[int64](5000),
	// 			Unit: to.Ptr("Count"),
	// 	}},
	// }
}
