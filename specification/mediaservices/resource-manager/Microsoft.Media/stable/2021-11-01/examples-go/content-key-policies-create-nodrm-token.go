package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-create-nodrm-token.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"contoso",
		"contosomedia",
		"PolicyWithClearKeyOptionAndSwtTokenRestriction",
		armmediaservices.ContentKeyPolicy{
			Properties: &armmediaservices.ContentKeyPolicyProperties{
				Description: to.Ptr("ArmPolicyDescription"),
				Options: []*armmediaservices.ContentKeyPolicyOption{
					{
						Name: to.Ptr("ClearKeyOption"),
						Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
						},
						Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
							Audience:  to.Ptr("urn:audience"),
							Issuer:    to.Ptr("urn:issuer"),
							PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
								ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
								KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
							},
							RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeSwt),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
