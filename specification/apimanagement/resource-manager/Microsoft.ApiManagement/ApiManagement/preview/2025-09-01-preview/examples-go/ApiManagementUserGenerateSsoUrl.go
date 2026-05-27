package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementUserGenerateSsoUrl.json
func ExampleUserClient_GenerateSsoURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().GenerateSsoURL(ctx, "rg1", "apimService1", "57127d485157a511ace86ae7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.UserClientGenerateSsoURLResponse{
	// 	GenerateSsoURLResult: armapimanagement.GenerateSsoURLResult{
	// 		Value: to.Ptr("https://apimService1.portal.azure-api.net/signin-sso?token=57127d485157a511ace86ae7%26201706051624%267VY18MlwAom***********2bYr2bDQHg21OzQsNakExQ%3d%3d"),
	// 	},
	// }
}
