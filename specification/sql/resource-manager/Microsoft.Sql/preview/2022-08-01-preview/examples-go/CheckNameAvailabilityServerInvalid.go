package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/CheckNameAvailabilityServerInvalid.json
func ExampleServersClient_CheckNameAvailability_checkForAServerNameThatIsInvalid() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().CheckNameAvailability(ctx, armsql.CheckNameAvailabilityRequest{
		Name: to.Ptr("SERVER1"),
		Type: to.Ptr("Microsoft.Sql/servers"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armsql.CheckNameAvailabilityResponse{
	// 	Name: to.Ptr("SERVER1"),
	// 	Available: to.Ptr(false),
	// 	Message: to.Ptr("Specified server name contains unsupported characters or is too long. Server name must be no longer than 63 characters long, contain only lower-case characters or digits, cannot contain '.' or '_' characters and can't start or end with '-' character."),
	// 	Reason: to.Ptr(armsql.CheckNameAvailabilityReasonInvalid),
	// }
}
