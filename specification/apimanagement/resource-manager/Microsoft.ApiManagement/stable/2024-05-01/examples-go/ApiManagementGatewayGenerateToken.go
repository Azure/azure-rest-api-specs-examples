package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGatewayGenerateToken.json
func ExampleGatewayClient_GenerateToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().GenerateToken(ctx, "rg1", "apimService1", "gw1", armapimanagement.GatewayTokenRequestContract{
		Expiry:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-21T00:44:24.284Z"); return t }()),
		KeyType: to.Ptr(armapimanagement.KeyTypePrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayTokenContract = armapimanagement.GatewayTokenContract{
	// 	Value: to.Ptr("gw1&201904210044&9A1GR1f5WIhFvFmzQG+xxxxxxxxxxx/kBeu87DWad3tkasUXuvPL+MgzlwUHyg=="),
	// }
}
