package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/EmailRegistrations_ActivateEmail.json
func ExampleEmailRegistrationsClient_ActivateEmail() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewEmailRegistrationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ActivateEmail(ctx,
		"East US 2",
		armdatashare.EmailRegistration{
			ActivationCode: to.Ptr("djsfhakj2lekowd3wepfklpwe9lpflcd"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
