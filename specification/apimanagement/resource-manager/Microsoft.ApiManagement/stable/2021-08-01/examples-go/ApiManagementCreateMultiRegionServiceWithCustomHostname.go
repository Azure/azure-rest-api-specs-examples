package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
func ExampleServiceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		armapimanagement.ServiceResource{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
				"tag3": to.Ptr("value3"),
			},
			Location: to.Ptr("West US"),
			Properties: &armapimanagement.ServiceProperties{
				AdditionalLocations: []*armapimanagement.AdditionalLocation{
					{
						DisableGateway: to.Ptr(true),
						Location:       to.Ptr("East US"),
						SKU: &armapimanagement.ServiceSKUProperties{
							Name:     to.Ptr(armapimanagement.SKUTypePremium),
							Capacity: to.Ptr[int32](1),
						},
					}},
				APIVersionConstraint: &armapimanagement.APIVersionConstraint{
					MinAPIVersion: to.Ptr("2019-01-01"),
				},
				HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
					{
						Type:                to.Ptr(armapimanagement.HostnameTypeProxy),
						CertificatePassword: to.Ptr("Password"),
						DefaultSSLBinding:   to.Ptr(true),
						EncodedCertificate:  to.Ptr("****** Base 64 Encoded Certificate ************"),
						HostName:            to.Ptr("gateway1.msitesting.net"),
					},
					{
						Type:                to.Ptr(armapimanagement.HostnameTypeManagement),
						CertificatePassword: to.Ptr("Password"),
						EncodedCertificate:  to.Ptr("****** Base 64 Encoded Certificate ************"),
						HostName:            to.Ptr("mgmt.msitesting.net"),
					},
					{
						Type:                to.Ptr(armapimanagement.HostnameTypePortal),
						CertificatePassword: to.Ptr("Password"),
						EncodedCertificate:  to.Ptr("****** Base 64 Encoded Certificate ************"),
						HostName:            to.Ptr("portal1.msitesting.net"),
					}},
				VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
				PublisherEmail:     to.Ptr("apim@autorestsdk.com"),
				PublisherName:      to.Ptr("autorestsdk"),
			},
			SKU: &armapimanagement.ServiceSKUProperties{
				Name:     to.Ptr(armapimanagement.SKUTypePremium),
				Capacity: to.Ptr[int32](1),
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
