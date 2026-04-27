package armappnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appnetwork/armappnetwork"
)

// Generated from example definition: 2025-08-01-preview/AppLinks_Update.json
func ExampleAppLinksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappnetwork.NewClientFactory("11809CA1-E126-4017-945E-AA795CD5C5A9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppLinksClient().BeginUpdate(ctx, "test_rg", "applink-test-01", armappnetwork.AppLinkUpdate{
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
			"cost-center": to.Ptr("platform"),
		},
	}, nil)
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
	// res = armappnetwork.AppLinksClientUpdateResponse{
	// 	AppLink: &armappnetwork.AppLink{
	// 		Properties: &armappnetwork.AppLinkProperties{
	// 			ProvisioningState: to.Ptr(armappnetwork.ProvisioningStateSucceeded),
	// 		},
	// 		Identity: &armappnetwork.ManagedServiceIdentity{
	// 			Type: to.Ptr(armappnetwork.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armappnetwork.UserAssignedIdentity{
	// 				"/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/test_rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/applink-identity": &armappnetwork.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("44444444-5555-6666-7777-888888888888"),
	// 					ClientID: to.Ptr("99999999-aaaa-bbbb-cccc-dddddddddddd"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"cost-center": to.Ptr("platform"),
	// 		},
	// 		Location: to.Ptr("westus2"),
	// 		ID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/test_rg/providers/Microsoft.AppLink/appLinks/applink-test-01"),
	// 		Name: to.Ptr("applink-test-01"),
	// 		Type: to.Ptr("Microsoft.AppLink/appLinks"),
	// 		SystemData: &armappnetwork.SystemData{
	// 			CreatedBy: to.Ptr("user01"),
	// 			CreatedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user02"),
	// 			LastModifiedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
	// 		},
	// 	},
	// }
}
