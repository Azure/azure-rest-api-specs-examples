package armfeatures_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/SubscriptionFeatureRegistrationPUT.json
func ExampleSubscriptionFeatureRegistrationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfeatures.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionFeatureRegistrationsClient().CreateOrUpdate(ctx, "subscriptionFeatureRegistrationGroupTestRG", "testFeature", &armfeatures.SubscriptionFeatureRegistrationsClientCreateOrUpdateOptions{SubscriptionFeatureRegistrationType: &armfeatures.SubscriptionFeatureRegistration{
		Properties: &armfeatures.SubscriptionFeatureRegistrationProperties{},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionFeatureRegistration = armfeatures.SubscriptionFeatureRegistration{
	// 	Name: to.Ptr("testFeature"),
	// 	Type: to.Ptr("Microsoft.Features/featureProviders/subscriptionFeatureRegistrations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Features/featureProviders/Microsoft.TestRP/subscriptionFeatureRegistrations/testFeature"),
	// 	Properties: &armfeatures.SubscriptionFeatureRegistrationProperties{
	// 		ApprovalType: to.Ptr(armfeatures.SubscriptionFeatureRegistrationApprovalTypeApprovalRequired),
	// 		AuthorizationProfile: &armfeatures.AuthorizationProfile{
	// 		},
	// 		FeatureName: to.Ptr("testFeature"),
	// 		ProviderNamespace: to.Ptr("Microsoft.TestRP"),
	// 		RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-26T01:57:51.734Z"); return t}()),
	// 		ReleaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-05T00:34:53.124Z"); return t}()),
	// 		State: to.Ptr(armfeatures.SubscriptionFeatureRegistrationStatePending),
	// 		SubscriptionID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 	},
	// }
}
