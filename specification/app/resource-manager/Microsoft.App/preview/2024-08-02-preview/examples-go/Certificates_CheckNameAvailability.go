package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Certificates_CheckNameAvailability.json
func ExampleNamespacesClient_CheckNameAvailability_certificatesCheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().CheckNameAvailability(ctx, "examplerg", "testcontainerenv", armappcontainers.CheckNameAvailabilityRequest{
		Name: to.Ptr("testcertificatename"),
		Type: to.Ptr("Microsoft.App/managedEnvironments/certificates"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armappcontainers.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr(""),
	// 	NameAvailable: to.Ptr(true),
	// 	Reason: to.Ptr(armappcontainers.CheckNameAvailabilityReason("None")),
	// }
}
