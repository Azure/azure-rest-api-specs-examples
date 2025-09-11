package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceTimeZoneGet.json
func ExampleTimeZonesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTimeZonesClient().Get(ctx, "canadaeast", "Haiti Standard Time", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TimeZone = armsql.TimeZone{
	// 	Name: to.Ptr("Haiti Standard Time"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
	// 	ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Haiti Standard Time"),
	// 	Properties: &armsql.TimeZoneProperties{
	// 		DisplayName: to.Ptr("(UTC-05:00) Haiti"),
	// 		TimeZoneID: to.Ptr("Haiti Standard Time"),
	// 	},
	// }
}
