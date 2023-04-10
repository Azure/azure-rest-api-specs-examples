package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9a65970ef1837c0af1800c906aa365ba95871b26/specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ContainerApps_ListCustomHostNameAnalysis.json
func ExampleContainerAppsClient_ListCustomHostNameAnalysis() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsClient().ListCustomHostNameAnalysis(ctx, "rg", "testcontainerApp0", &armappcontainers.ContainerAppsClientListCustomHostNameAnalysisOptions{CustomHostname: to.Ptr("my.name.corp")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomHostnameAnalysisResult = armappcontainers.CustomHostnameAnalysisResult{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerApp0"),
	// 	Properties: &armappcontainers.CustomHostnameAnalysisResultProperties{
	// 		ARecords: []*string{
	// 			to.Ptr("aRecord1"),
	// 			to.Ptr("aRecord2")},
	// 			AlternateCNameRecords: []*string{
	// 				to.Ptr("cNameRecord1"),
	// 				to.Ptr("cNameRecord2")},
	// 				AlternateTxtRecords: []*string{
	// 					to.Ptr("txtRecord1"),
	// 					to.Ptr("txtRecord2")},
	// 					CNameRecords: []*string{
	// 						to.Ptr("cNameRecord1"),
	// 						to.Ptr("cNameRecord2")},
	// 						ConflictingContainerAppResourceID: to.Ptr(""),
	// 						CustomDomainVerificationFailureInfo: &armappcontainers.DefaultErrorResponse{
	// 						},
	// 						CustomDomainVerificationTest: to.Ptr(armappcontainers.DNSVerificationTestResultPassed),
	// 						HasConflictOnManagedEnvironment: to.Ptr(false),
	// 						HostName: to.Ptr("my.name.corp"),
	// 						IsHostnameAlreadyVerified: to.Ptr(true),
	// 						TxtRecords: []*string{
	// 							to.Ptr("txtRecord1"),
	// 							to.Ptr("txtRecord2")},
	// 						},
	// 					}
}
