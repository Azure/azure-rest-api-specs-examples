package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/PutRegistrationAssignment.json
func ExampleRegistrationAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistrationAssignmentsClient().BeginCreateOrUpdate(ctx, "subscription/0afefe50-734e-4610-8a82-a144ahf49dea", "26c128c2-fefa-4340-9bb1-6e081c90ada2", armmanagedservices.RegistrationAssignment{
		Properties: &armmanagedservices.RegistrationAssignmentProperties{
			RegistrationDefinitionID: to.Ptr("/subscriptions/0afefe50-734e-4610-8a82-a144ahf49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-6e081c90ada2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegistrationAssignment = armmanagedservices.RegistrationAssignment{
	// 	Name: to.Ptr("484a7d5f-9729-4b87-bc9b-26610985a013"),
	// 	Type: to.Ptr("Microsoft.ManagedServices/registrationAssignments"),
	// 	ID: to.Ptr("/subscriptions/0afefe50-734e-4610-8c82-a144aff49dea/providers/Microsoft.ManagedServices/registrationAssignments/484a7d5f-9729-4b87-bc9b-26610985a013"),
	// 	Properties: &armmanagedservices.RegistrationAssignmentProperties{
	// 		ProvisioningState: to.Ptr(armmanagedservices.ProvisioningStateSucceeded),
	// 		RegistrationDefinitionID: to.Ptr("/subscriptions/0afefe50-734e-4610-8c82-a144aff49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-8e081c90ada2"),
	// 	},
	// }
}
