package armmongocluster_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster/TempTypeSpecFiles/DocumentDB.MongoCluster.Management/examples/2024-07-01/MongoClusters_NameAvailability_AlreadyExists.json
func ExampleMongoClustersClient_CheckNameAvailability_checksAndReturnsThatTheMongoClusterNameIsAlreadyInUse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMongoClustersClient().CheckNameAvailability(ctx, "westus2", armmongocluster.CheckNameAvailabilityRequest{
		Name: to.Ptr("existingmongocluster"),
		Type: to.Ptr("Microsoft.DocumentDB/mongoClusters"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongocluster.MongoClustersClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResponse: &armmongocluster.CheckNameAvailabilityResponse{
	// 		NameAvailable: to.Ptr(false),
	// 		Reason: to.Ptr(armmongocluster.CheckNameAvailabilityReasonAlreadyExists),
	// 		Message: to.Ptr("Cluster name 'existingmongocluster' is already in use."),
	// 	},
	// }
}
