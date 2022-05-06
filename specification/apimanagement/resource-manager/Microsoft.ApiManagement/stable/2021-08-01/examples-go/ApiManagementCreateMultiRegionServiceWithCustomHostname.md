Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv0.5.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
func ExampleServiceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewServiceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.ServiceResource{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
				"tag3": to.Ptr("value3"),
			},
			Location: to.Ptr("<location>"),
			Properties: &armapimanagement.ServiceProperties{
				AdditionalLocations: []*armapimanagement.AdditionalLocation{
					{
						DisableGateway: to.Ptr(true),
						Location:       to.Ptr("<location>"),
						SKU: &armapimanagement.ServiceSKUProperties{
							Name:     to.Ptr(armapimanagement.SKUTypePremium),
							Capacity: to.Ptr[int32](1),
						},
					}},
				APIVersionConstraint: &armapimanagement.APIVersionConstraint{
					MinAPIVersion: to.Ptr("<min-apiversion>"),
				},
				HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
					{
						Type:                to.Ptr(armapimanagement.HostnameTypeProxy),
						CertificatePassword: to.Ptr("<certificate-password>"),
						DefaultSSLBinding:   to.Ptr(true),
						EncodedCertificate:  to.Ptr("<encoded-certificate>"),
						HostName:            to.Ptr("<host-name>"),
					},
					{
						Type:                to.Ptr(armapimanagement.HostnameTypeManagement),
						CertificatePassword: to.Ptr("<certificate-password>"),
						EncodedCertificate:  to.Ptr("<encoded-certificate>"),
						HostName:            to.Ptr("<host-name>"),
					},
					{
						Type:                to.Ptr(armapimanagement.HostnameTypePortal),
						CertificatePassword: to.Ptr("<certificate-password>"),
						EncodedCertificate:  to.Ptr("<encoded-certificate>"),
						HostName:            to.Ptr("<host-name>"),
					}},
				VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
				PublisherEmail:     to.Ptr("<publisher-email>"),
				PublisherName:      to.Ptr("<publisher-name>"),
			},
			SKU: &armapimanagement.ServiceSKUProperties{
				Name:     to.Ptr(armapimanagement.SKUTypePremium),
				Capacity: to.Ptr[int32](1),
			},
		},
		&armapimanagement.ServiceClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
