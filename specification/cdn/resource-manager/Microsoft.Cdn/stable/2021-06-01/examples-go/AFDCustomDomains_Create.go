package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_Create.json
func ExampleAFDCustomDomainsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewAFDCustomDomainsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"RG",
		"profile1",
		"domain1",
		armcdn.AFDDomain{
			Properties: &armcdn.AFDDomainProperties{
				AzureDNSZone: &armcdn.ResourceReference{
					ID: to.Ptr(""),
				},
				TLSSettings: &armcdn.AFDDomainHTTPSParameters{
					CertificateType:   to.Ptr(armcdn.AfdCertificateTypeManagedCertificate),
					MinimumTLSVersion: to.Ptr(armcdn.AfdMinimumTLSVersionTLS12),
				},
				HostName: to.Ptr("www.someDomain.net"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
