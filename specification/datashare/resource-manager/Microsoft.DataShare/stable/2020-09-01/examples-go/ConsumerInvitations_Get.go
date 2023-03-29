package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ConsumerInvitations_Get.json
func ExampleConsumerInvitationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConsumerInvitationsClient().Get(ctx, "East US 2", "dfbbc788-19eb-4607-a5a1-c74181bfff03", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConsumerInvitation = armdatashare.ConsumerInvitation{
	// 	Name: to.Ptr("invitation1"),
	// 	Type: to.Ptr("Microsoft.DataShare/locations/consumerInvitations"),
	// 	ID: to.Ptr("providers/Microsoft.DataShare/locations/eastus2/consumerInvitations/4256e2cf-0f82-4865-961b-12f83333f487"),
	// 	Properties: &armdatashare.ConsumerInvitationProperties{
	// 		Description: to.Ptr("Some share"),
	// 		DataSetCount: to.Ptr[int32](1),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-26T22:33:24.5785265Z"); return t}()),
	// 		InvitationID: to.Ptr("4256e2cf-0f82-4865-961b-12f83333f487"),
	// 		InvitationStatus: to.Ptr(armdatashare.InvitationStatusAccepted),
	// 		Location: to.Ptr("eastus2"),
	// 		ProviderEmail: to.Ptr("john.smith@microsoft.com"),
	// 		ProviderName: to.Ptr("John Smith"),
	// 		ProviderTenantName: to.Ptr("Microsoft"),
	// 		RespondedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-15T02:01:51.8953054Z"); return t}()),
	// 		SentAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-15T01:45:25.6226946Z"); return t}()),
	// 		ShareName: to.Ptr("share1"),
	// 		TermsOfUse: to.Ptr("Confidential"),
	// 	},
	// }
}
