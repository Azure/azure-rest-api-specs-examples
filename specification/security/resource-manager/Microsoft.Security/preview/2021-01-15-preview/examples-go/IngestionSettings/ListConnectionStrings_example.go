package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListConnectionStrings_example.json
func ExampleIngestionSettingsClient_ListConnectionStrings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIngestionSettingsClient().ListConnectionStrings(ctx, "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionStrings = armsecurity.ConnectionStrings{
	// 	Value: []*armsecurity.IngestionConnectionString{
	// 		{
	// 			Location: to.Ptr("CUS"),
	// 			Value: to.Ptr("<connection string>"),
	// 		},
	// 		{
	// 			Location: to.Ptr("WEU"),
	// 			Value: to.Ptr("<connection string>"),
	// 	}},
	// }
}
