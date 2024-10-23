package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/ValidateLinkerSuccess.json
func ExampleLinkerClient_BeginValidate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLinkerClient().BeginValidate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "linkName", nil)
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
	// res.ValidateOperationResult = armservicelinker.ValidateOperationResult{
	// 	Properties: &armservicelinker.ValidateResult{
	// 		AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 		IsConnectionAvailable: to.Ptr(true),
	// 		ReportEndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:06:09.000Z"); return t}()),
	// 		ReportStartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
	// 		SourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 		TargetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 		ValidationDetail: []*armservicelinker.ValidationResultItem{
	// 			{
	// 				Name: to.Ptr("TargetExistence"),
	// 				Description: to.Ptr("The target existence is validated"),
	// 				Result: to.Ptr(armservicelinker.ValidationResultStatusSuccess),
	// 			},
	// 			{
	// 				Name: to.Ptr("TargetNetworkAccess"),
	// 				Description: to.Ptr("Deny public network access is set to yes. Please confirm you are using private endpoint connection to access target resource."),
	// 				Result: to.Ptr(armservicelinker.ValidationResultStatusWarning),
	// 		}},
	// 	},
	// }
}
