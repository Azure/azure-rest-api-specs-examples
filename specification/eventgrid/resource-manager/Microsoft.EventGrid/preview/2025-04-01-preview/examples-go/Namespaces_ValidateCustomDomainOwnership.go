package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_ValidateCustomDomainOwnership.json
func ExampleNamespacesClient_BeginValidateCustomDomainOwnership() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginValidateCustomDomainOwnership(ctx, "examplerg", "exampleNamespaceName1", nil)
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
	// res.CustomDomainOwnershipValidationResult = armeventgrid.CustomDomainOwnershipValidationResult{
	// 	CustomDomainsForTopicSpacesConfiguration: []*armeventgrid.CustomDomainConfiguration{
	// 		{
	// 			CertificateURL: to.Ptr("string"),
	// 			ExpectedTxtRecordName: to.Ptr("txt-record-name-2"),
	// 			ExpectedTxtRecordValue: to.Ptr("txt-record-value-2"),
	// 			FullyQualifiedDomainName: to.Ptr("custom-domain-name-2"),
	// 			Identity: &armeventgrid.CustomDomainIdentity{
	// 				Type: to.Ptr(armeventgrid.CustomDomainIdentityTypeSystemAssigned),
	// 			},
	// 			ValidationState: to.Ptr(armeventgrid.CustomDomainValidationStatePending),
	// 	}},
	// 	CustomDomainsForTopicsConfiguration: []*armeventgrid.CustomDomainConfiguration{
	// 		{
	// 			CertificateURL: to.Ptr("string"),
	// 			ExpectedTxtRecordName: to.Ptr("txt-record-name-1"),
	// 			ExpectedTxtRecordValue: to.Ptr("txt-record-value-1"),
	// 			FullyQualifiedDomainName: to.Ptr("custom-domain-name-1"),
	// 			Identity: &armeventgrid.CustomDomainIdentity{
	// 				Type: to.Ptr(armeventgrid.CustomDomainIdentityTypeSystemAssigned),
	// 			},
	// 			ValidationState: to.Ptr(armeventgrid.CustomDomainValidationStatePending),
	// 	}},
	// }
}
