package armmanagedservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// x-ms-original-file: specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2020-02-01-preview/examples/DeleteRegistrationAssignment.json
func ExampleRegistrationAssignmentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedservices.NewRegistrationAssignmentsClient(cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<scope>",
		"<registration-assignment-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
