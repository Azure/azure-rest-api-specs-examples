package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-01/UpdateAccountClearTrustedCertificateSources.json
func ExampleAccountsClient_BeginUpdate_clearAccountPrivateCaTrustUsingAnEmptyArray() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginUpdate(ctx, "myResourceGroup", "myAccount", armcognitiveservices.Account{
		Properties: &armcognitiveservices.AccountProperties{
			TrustedCertificateSources: []*armcognitiveservices.TrustedCertificateSource{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AccountsClientUpdateResponse{
	// 	Account: armcognitiveservices.Account{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount"),
	// 		Name: to.Ptr("myAccount"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 		Kind: to.Ptr("AIServices"),
	// 		Location: to.Ptr("eastus"),
	// 		SKU: &armcognitiveservices.SKU{
	// 			Name: to.Ptr("S0"),
	// 		},
	// 		Properties: &armcognitiveservices.AccountProperties{
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			AllowProjectManagement: to.Ptr(true),
	// 			TrustedCertificateSources: []*armcognitiveservices.TrustedCertificateSource{
	// 			},
	// 		},
	// 	},
	// }
}
