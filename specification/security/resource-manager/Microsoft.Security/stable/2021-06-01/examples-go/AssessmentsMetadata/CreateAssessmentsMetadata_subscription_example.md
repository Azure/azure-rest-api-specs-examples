Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	res, err := client.CreateInSubscription(ctx,
		"<assessment-metadata-name>",
		armsecurity.AssessmentMetadataResponse{
			Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
				Description:    to.StringPtr("<description>"),
				AssessmentType: armsecurity.AssessmentType("CustomerManaged").ToPtr(),
				Categories: []*armsecurity.Categories{
					armsecurity.Categories("Compute").ToPtr()},
				DisplayName:            to.StringPtr("<display-name>"),
				ImplementationEffort:   armsecurity.ImplementationEffort("Low").ToPtr(),
				RemediationDescription: to.StringPtr("<remediation-description>"),
				Severity:               armsecurity.Severity("Medium").ToPtr(),
				Threats: []*armsecurity.Threats{
					armsecurity.Threats("dataExfiltration").ToPtr(),
					armsecurity.Threats("dataSpillage").ToPtr(),
					armsecurity.Threats("maliciousInsider").ToPtr()},
				UserImpact: armsecurity.UserImpact("Low").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssessmentsMetadataClientCreateInSubscriptionResult)
}
```
