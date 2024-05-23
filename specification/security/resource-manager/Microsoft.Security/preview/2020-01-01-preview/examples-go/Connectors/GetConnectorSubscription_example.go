package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetConnectorSubscription_example.json
func ExampleAccountConnectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountConnectorsClient().Get(ctx, "aws_dev1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectorSetting = armsecurity.ConnectorSetting{
	// 	Name: to.Ptr("aws_dev1"),
	// 	Type: to.Ptr("Microsoft.Security/connectors"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/connectors/aws_dev1"),
	// 	Properties: &armsecurity.ConnectorSettingProperties{
	// 		AuthenticationDetails: &armsecurity.AwsCredsAuthenticationDetailsProperties{
	// 			AuthenticationProvisioningState: to.Ptr(armsecurity.AuthenticationProvisioningStateValid),
	// 			AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsCreds),
	// 			GrantedPermissions: []*armsecurity.PermissionProperty{
	// 				to.Ptr(armsecurity.PermissionPropertyAWSAWSSecurityHubReadOnlyAccess),
	// 				to.Ptr(armsecurity.PermissionPropertyAWSSecurityAudit),
	// 				to.Ptr(armsecurity.PermissionPropertyAWSAmazonSSMAutomationRole)},
	// 				AccountID: to.Ptr("922315681122"),
	// 				AwsAccessKeyID: to.Ptr(""),
	// 				AwsSecretAccessKey: to.Ptr(""),
	// 			},
	// 			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
	// 				AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
	// 				HybridComputeProvisioningState: to.Ptr(armsecurity.HybridComputeProvisioningStateValid),
	// 				ProxyServer: &armsecurity.ProxyServerProperties{
	// 					IP: to.Ptr("287.221.107.152"),
	// 					Port: to.Ptr("34"),
	// 				},
	// 				Region: to.Ptr("West US 2"),
	// 				ResourceGroupName: to.Ptr("AwsConnectorRG"),
	// 				ServicePrincipal: &armsecurity.ServicePrincipalProperties{
	// 					ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
	// 				},
	// 			},
	// 		},
	// 	}
}
