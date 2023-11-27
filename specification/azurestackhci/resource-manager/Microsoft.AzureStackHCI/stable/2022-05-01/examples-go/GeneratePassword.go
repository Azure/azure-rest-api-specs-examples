package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/GeneratePassword.json
func ExampleArcSettingsClient_GeneratePassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArcSettingsClient().GeneratePassword(ctx, "test-rg", "myCluster", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PasswordCredential = armazurestackhci.PasswordCredential{
	// 	EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2121-02-17T10:54:23.626Z"); return t}()),
	// 	KeyID: to.Ptr("00000000-2d47-4fb2-8ed2-fed71a5c197b"),
	// 	SecretText: to.Ptr("secretText"),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-17T10:54:23.626Z"); return t}()),
	// }
}
