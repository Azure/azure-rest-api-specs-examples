package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ContainerApps_CreateOrUpdate.json
func ExampleContainerAppsClient_BeginCreateOrUpdate_createOrUpdateContainerApp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewContainerAppsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg", "testcontainerApp0", armappcontainers.ContainerApp{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ContainerAppProperties{
			Configuration: &armappcontainers.Configuration{
				Dapr: &armappcontainers.Dapr{
					AppPort:            to.Ptr[int32](3000),
					AppProtocol:        to.Ptr(armappcontainers.AppProtocolHTTP),
					EnableAPILogging:   to.Ptr(true),
					Enabled:            to.Ptr(true),
					HTTPMaxRequestSize: to.Ptr[int32](10),
					HTTPReadBufferSize: to.Ptr[int32](30),
					LogLevel:           to.Ptr(armappcontainers.LogLevelDebug),
				},
				Ingress: &armappcontainers.Ingress{
					CustomDomains: []*armappcontainers.CustomDomain{
						{
							Name:          to.Ptr("www.my-name.com"),
							BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
							CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
						},
						{
							Name:          to.Ptr("www.my-other-name.com"),
							BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
							CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
						}},
					External: to.Ptr(true),
					IPSecurityRestrictions: []*armappcontainers.IPSecurityRestrictionRule{
						{
							Name:           to.Ptr("Allow work IP A subnet"),
							Description:    to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
							Action:         to.Ptr(armappcontainers.ActionAllow),
							IPAddressRange: to.Ptr("192.168.1.1/32"),
						},
						{
							Name:           to.Ptr("Allow work IP B subnet"),
							Description:    to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
							Action:         to.Ptr(armappcontainers.ActionAllow),
							IPAddressRange: to.Ptr("192.168.1.1/8"),
						}},
					TargetPort: to.Ptr[int32](3000),
					Traffic: []*armappcontainers.TrafficWeight{
						{
							Label:        to.Ptr("production"),
							RevisionName: to.Ptr("testcontainerApp0-ab1234"),
							Weight:       to.Ptr[int32](100),
						}},
				},
				MaxInactiveRevisions: to.Ptr[int32](10),
			},
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			Template: &armappcontainers.Template{
				Containers: []*armappcontainers.Container{
					{
						Name:  to.Ptr("testcontainerApp0"),
						Image: to.Ptr("repo/testcontainerApp0:v1"),
						Probes: []*armappcontainers.ContainerAppProbe{
							{
								Type: to.Ptr(armappcontainers.TypeLiveness),
								HTTPGet: &armappcontainers.ContainerAppProbeHTTPGet{
									Path: to.Ptr("/health"),
									HTTPHeaders: []*armappcontainers.ContainerAppProbeHTTPGetHTTPHeadersItem{
										{
											Name:  to.Ptr("Custom-Header"),
											Value: to.Ptr("Awesome"),
										}},
									Port: to.Ptr[int32](8080),
								},
								InitialDelaySeconds: to.Ptr[int32](3),
								PeriodSeconds:       to.Ptr[int32](3),
							}},
					}},
				InitContainers: []*armappcontainers.InitContainer{
					{
						Name: to.Ptr("testinitcontainerApp0"),
						Args: []*string{
							to.Ptr("-c"),
							to.Ptr("while true; do echo hello; sleep 10;done")},
						Command: []*string{
							to.Ptr("/bin/sh")},
						Image: to.Ptr("repo/testcontainerApp0:v4"),
						Resources: &armappcontainers.ContainerResources{
							CPU:    to.Ptr[float64](0.2),
							Memory: to.Ptr("100Mi"),
						},
					}},
				Scale: &armappcontainers.Scale{
					MaxReplicas: to.Ptr[int32](5),
					MinReplicas: to.Ptr[int32](1),
					Rules: []*armappcontainers.ScaleRule{
						{
							Name: to.Ptr("httpscalingrule"),
							Custom: &armappcontainers.CustomScaleRule{
								Type: to.Ptr("http"),
								Metadata: map[string]*string{
									"concurrentRequests": to.Ptr("50"),
								},
							},
						}},
				},
			},
			WorkloadProfileType: to.Ptr("GeneralPurpose"),
		},
	}, nil)
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
