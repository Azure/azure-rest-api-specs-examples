package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsCredConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate_awsCredCreateACloudAccountConnectorForASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "aws_dev1", armsecurity.ConnectorSetting{
		Properties: &armsecurity.ConnectorSettingProperties{
			AuthenticationDetails: &armsecurity.AwsCredsAuthenticationDetailsProperties{
				AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsCreds),
				AwsAccessKeyID:     to.Ptr("AKIARPZCNODDNAEQFSOE"),
				AwsSecretAccessKey: to.Ptr("<awsSecretAccessKey>"),
			},
			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
				AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
				ProxyServer: &armsecurity.ProxyServerProperties{
					IP:   to.Ptr("167.220.197.140"),
					Port: to.Ptr("34"),
				},
				Region:            to.Ptr("West US 2"),
				ResourceGroupName: to.Ptr("AwsConnectorRG"),
				ServicePrincipal: &armsecurity.ServicePrincipalProperties{
					ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
					Secret:        to.Ptr("<secret>"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
