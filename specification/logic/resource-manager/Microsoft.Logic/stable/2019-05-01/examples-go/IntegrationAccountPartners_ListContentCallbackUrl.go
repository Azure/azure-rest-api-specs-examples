package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_ListContentCallbackUrl.json
func ExampleIntegrationAccountPartnersClient_ListContentCallbackURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountPartnersClient().ListContentCallbackURL(ctx, "testResourceGroup", "testIntegrationAccount", "testPartner", armlogic.GetCallbackURLParameters{
		KeyType:  to.Ptr(armlogic.KeyTypePrimary),
		NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-19T16:00:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerCallbackURL = armlogic.WorkflowTriggerCallbackURL{
	// 	Method: to.Ptr("GET"),
	// 	BasePath: to.Ptr("https://prod-00.westus.logic.azure.com/integrationAccounts/0fdabc3a76514ca48dede71c73d9fe97/partners/testPartner/contents/Value"),
	// 	Queries: &armlogic.WorkflowTriggerListCallbackURLQueries{
	// 		APIVersion: to.Ptr("2015-08-01-preview"),
	// 	},
	// 	Value: to.Ptr("https://prod-00.westus.logic.azure.com:443/integrationAccounts/0fdabc3a76514ca48dede71c73d9fe97/partners/testPartner/contents/Value?api-version=2015-08-01-preview&sp=%2Fpartners%2FtestPartner%2Fread&sv=1.0&sig=VK_mbQPTHTa3ezhsrI8IctckwjlL3GdJmroQH_baYj4"),
	// }
}
