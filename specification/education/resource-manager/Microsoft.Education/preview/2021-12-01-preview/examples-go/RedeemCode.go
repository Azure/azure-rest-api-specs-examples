package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/RedeemCode.json
func ExampleManagementClient_RedeemInvitationCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeducation.NewManagementClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.RedeemInvitationCode(ctx,
		armeducation.RedeemRequest{
			FirstName:  to.Ptr("test"),
			LastName:   to.Ptr("user"),
			RedeemCode: to.Ptr("exampleRedeemCode"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
