package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBDataTransferServiceCreate.json
func ExampleServiceClient_BeginCreate_dataTransferServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceClient().BeginCreate(ctx, "rg1", "ddb1", "DataTransfer", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.DataTransferServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeDataTransfer),
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
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("DataTransfer"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/DataTransfer"),
	// 	Properties: &armcosmos.DataTransferServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.462Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeDataTransfer),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.DataTransferRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("DataTransfer-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		}},
	// 	},
	// }
}
