package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/OperationResultsGet.json
func ExampleOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationResultsClient().Get(ctx, "westus", "exampleid", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationResultsDescription = armm365securityandcompliance.OperationResultsDescription{
	// 	Name: to.Ptr("servicename"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.M365SecurityAndCompliance/locations/westus/operationresults/exampleid"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("2020-01-11T06:03:30.2716301Z"),
	// 	Status: to.Ptr(armm365securityandcompliance.OperationResultStatusRequested),
	// }
}
