package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_CreateOrUpdate.json
func ExampleIntegrationAccountPartnersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationAccountPartnersClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testResourceGroup",
		"testIntegrationAccount",
		"testPartner",
		armlogic.IntegrationAccountPartner{
			Location: to.Ptr("westus"),
			Tags:     map[string]*string{},
			Properties: &armlogic.IntegrationAccountPartnerProperties{
				Content: &armlogic.PartnerContent{
					B2B: &armlogic.B2BPartnerContent{
						BusinessIdentities: []*armlogic.BusinessIdentity{
							{
								Qualifier: to.Ptr("AA"),
								Value:     to.Ptr("ZZ"),
							}},
					},
				},
				Metadata:    map[string]interface{}{},
				PartnerType: to.Ptr(armlogic.PartnerTypeB2B),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
