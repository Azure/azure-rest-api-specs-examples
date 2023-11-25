package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_CreateOrUpdate.json
func ExampleIntegrationAccountPartnersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountPartnersClient().CreateOrUpdate(ctx, "testResourceGroup", "testIntegrationAccount", "testPartner", armlogic.IntegrationAccountPartner{
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
			Metadata:    map[string]any{},
			PartnerType: to.Ptr(armlogic.PartnerTypeB2B),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountPartner = armlogic.IntegrationAccountPartner{
	// 	Name: to.Ptr("testPartner"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/partners"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/flowrg/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/partners/testPartner"),
	// 	Properties: &armlogic.IntegrationAccountPartnerProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T20:06:23.590Z"); return t}()),
	// 		Content: &armlogic.PartnerContent{
	// 			B2B: &armlogic.B2BPartnerContent{
	// 				BusinessIdentities: []*armlogic.BusinessIdentity{
	// 					{
	// 						Qualifier: to.Ptr("AA"),
	// 						Value: to.Ptr("ZZ"),
	// 				}},
	// 			},
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T20:06:23.589Z"); return t}()),
	// 		Metadata: map[string]any{
	// 		},
	// 		PartnerType: to.Ptr(armlogic.PartnerTypeB2B),
	// 	},
	// }
}
