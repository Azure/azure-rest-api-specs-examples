package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecommendedActionSessionOperationStatus.json
func ExampleLocationBasedRecommendedActionSessionsOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationBasedRecommendedActionSessionsOperationStatusClient().Get(ctx, "WestUS", "aaaabbbb-cccc-dddd-0000-111122223333", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecommendedActionSessionsOperationStatus = armmariadb.RecommendedActionSessionsOperationStatus{
	// 	Name: to.Ptr("aaaabbbb-cccc-dddd-0000-111122223333"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T17:15:00.000Z"); return t}()),
	// 	Status: to.Ptr("succeeded"),
	// }
}
