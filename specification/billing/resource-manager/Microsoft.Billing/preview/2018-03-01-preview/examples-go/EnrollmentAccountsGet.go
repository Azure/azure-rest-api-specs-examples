package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsGet.json
func ExampleEnrollmentAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnrollmentAccountsClient().Get(ctx, "e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnrollmentAccountSummary = armbilling.EnrollmentAccountSummary{
	// 	Name: to.Ptr("e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556"),
	// 	Type: to.Ptr("Microsoft.Billing/enrollmentAccounts"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/enrollmentAccounts/e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556"),
	// 	Properties: &armbilling.EnrollmentAccountSummaryProperties{
	// 		PrincipalName: to.Ptr("kathy@contoso.com"),
	// 	},
	// }
}
