package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/mongo-cluster/CosmosDBMongoClusterNameAvailability_AlreadyExists.json
func ExampleMongoClustersClient_CheckNameAvailability_checkNameAvailabilityAlreadyExistsResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMongoClustersClient().CheckNameAvailability(ctx, "westus2", armcosmos.CheckNameAvailabilityRequest{
		Name: to.Ptr("existingmongocluster"),
		Type: to.Ptr("Microsoft.DocumentDB/mongoClusters"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armcosmos.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr("Cluster name 'existingmongocluster' is already in use."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armcosmos.CheckNameAvailabilityReasonAlreadyExists),
	// }
}
