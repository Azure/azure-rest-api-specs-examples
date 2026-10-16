package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/GeneratePassword.json
func ExampleArcSettingsClient_GeneratePassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
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
	// res = armazurestackhci.ArcSettingsClientGeneratePasswordResponse{
	// 	PasswordCredential: armazurestackhci.PasswordCredential{
	// 		EndDateTime: to.Ptr(time.Date(2121, time.February, 17, 10, 54, 23, 626437700, time.UTC)),
	// 		KeyID: to.Ptr("00000000-2d47-4fb2-8ed2-fed71a5c197b"),
	// 		SecretText: to.Ptr("secretText"),
	// 		StartDateTime: to.Ptr(time.Date(2022, time.February, 17, 10, 54, 23, 626400500, time.UTC)),
	// 	},
	// }
}
