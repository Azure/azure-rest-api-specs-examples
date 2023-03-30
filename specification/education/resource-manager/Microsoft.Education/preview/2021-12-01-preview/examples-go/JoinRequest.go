package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9b91929c304f8fb44002267b6c98d9fb9dde014/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequest.json
func ExampleJoinRequestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJoinRequestsClient().Get(ctx, "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", "{joinRequestName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JoinRequestDetails = armeducation.JoinRequestDetails{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Education/JoinRequest"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests/{joinRequestName}"),
	// 	Properties: &armeducation.JoinRequestProperties{
	// 		Email: to.Ptr("test@contoso.com"),
	// 		FirstName: to.Ptr("test"),
	// 		LastName: to.Ptr("user"),
	// 		Status: to.Ptr(armeducation.JoinRequestStatusPending),
	// 	},
	// }
}
