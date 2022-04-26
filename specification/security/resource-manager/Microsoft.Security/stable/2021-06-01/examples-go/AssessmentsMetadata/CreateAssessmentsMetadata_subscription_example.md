Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateInSubscription(ctx,
		"<assessment-metadata-name>",
		armsecurity.AssessmentMetadataResponse{
			Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
				Description:    to.Ptr("<description>"),
				AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomerManaged),
				Categories: []*armsecurity.Categories{
					to.Ptr(armsecurity.CategoriesCompute)},
				DisplayName:            to.Ptr("<display-name>"),
				ImplementationEffort:   to.Ptr(armsecurity.ImplementationEffortLow),
				RemediationDescription: to.Ptr("<remediation-description>"),
				Severity:               to.Ptr(armsecurity.SeverityMedium),
				Threats: []*armsecurity.Threats{
					to.Ptr(armsecurity.ThreatsDataExfiltration),
					to.Ptr(armsecurity.ThreatsDataSpillage),
					to.Ptr(armsecurity.ThreatsMaliciousInsider)},
				UserImpact: to.Ptr(armsecurity.UserImpactLow),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
