package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBDataTransferServiceGet.json
func ExampleServiceClient_Get_dataTransferServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().Get(ctx, "rg1", "ddb1", "DataTransfer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
