package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/RegistryGenerateCredentials.json
func ExampleRegistriesClient_BeginGenerateCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistriesClient().BeginGenerateCredentials(ctx, "myResourceGroup", "myRegistry", armcontainerregistry.GenerateCredentialsParameters{
		Expiry:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T15:59:59.070Z"); return t }()),
		TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/myToken"),
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
	// res.GenerateCredentialsResult = armcontainerregistry.GenerateCredentialsResult{
	// 	Passwords: []*armcontainerregistry.TokenPassword{
	// 		{
	// 			Name: to.Ptr(armcontainerregistry.TokenPasswordNamePassword1),
	// 			Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T15:59:59.070Z"); return t}()),
	// 			Value: to.Ptr("00000000000000000000000000000000"),
	// 		},
	// 		{
	// 			Name: to.Ptr(armcontainerregistry.TokenPasswordNamePassword2),
	// 			Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T15:59:59.070Z"); return t}()),
	// 			Value: to.Ptr("00000000000000000000000000000000"),
	// 	}},
	// 	Username: to.Ptr("myToken"),
	// }
}
