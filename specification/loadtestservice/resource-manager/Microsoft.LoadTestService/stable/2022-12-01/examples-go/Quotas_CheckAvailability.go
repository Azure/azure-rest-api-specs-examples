package armloadtestservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtestservice/armloadtestservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_CheckAvailability.json
func ExampleQuotasClient_CheckAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armloadtestservice.NewQuotasClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckAvailability(ctx, "westus", "testQuotaBucket", armloadtestservice.QuotaBucketRequest{
		Properties: &armloadtestservice.QuotaBucketRequestProperties{
			CurrentQuota: to.Ptr[int32](40),
			CurrentUsage: to.Ptr[int32](20),
			Dimensions: &armloadtestservice.QuotaBucketRequestPropertiesDimensions{
				Location:       to.Ptr("westus"),
				SubscriptionID: to.Ptr("testsubscriptionId"),
			},
			NewQuota: to.Ptr[int32](50),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
