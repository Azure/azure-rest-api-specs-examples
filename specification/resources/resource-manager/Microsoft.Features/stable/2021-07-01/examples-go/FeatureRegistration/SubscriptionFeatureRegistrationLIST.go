package armfeatures_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/SubscriptionFeatureRegistrationLIST.json
func ExampleSubscriptionFeatureRegistrationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfeatures.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionFeatureRegistrationsClient().NewListBySubscriptionPager("subscriptionFeatureRegistrationGroupTestRG", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SubscriptionFeatureRegistrationList = armfeatures.SubscriptionFeatureRegistrationList{
		// 	Value: []*armfeatures.SubscriptionFeatureRegistration{
		// 		{
		// 			Name: to.Ptr("testFeature"),
		// 			Type: to.Ptr("Microsoft.Features/featureProviders/subscriptionFeatureRegistrations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Features/featureProviders/Microsoft.TestRP/subscriptionFeatureRegistrations/testFeature"),
		// 			Properties: &armfeatures.SubscriptionFeatureRegistrationProperties{
		// 				ApprovalType: to.Ptr(armfeatures.SubscriptionFeatureRegistrationApprovalTypeApprovalRequired),
		// 				AuthorizationProfile: &armfeatures.AuthorizationProfile{
		// 				},
		// 				FeatureName: to.Ptr("testFeature"),
		// 				ProviderNamespace: to.Ptr("Microsoft.TestRP"),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-26T01:57:51.734Z"); return t}()),
		// 				ReleaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-05T00:34:53.124Z"); return t}()),
		// 				State: to.Ptr(armfeatures.SubscriptionFeatureRegistrationStatePending),
		// 				SubscriptionID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 			},
		// 	}},
		// }
	}
}
