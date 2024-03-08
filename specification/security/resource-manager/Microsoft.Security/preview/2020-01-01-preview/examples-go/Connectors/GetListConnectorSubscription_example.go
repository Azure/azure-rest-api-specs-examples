package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetListConnectorSubscription_example.json
func ExampleAccountConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountConnectorsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ConnectorSettingList = armsecurity.ConnectorSettingList{
		// 	Value: []*armsecurity.ConnectorSetting{
		// 		{
		// 			Name: to.Ptr("aws_dev1"),
		// 			Type: to.Ptr("Microsoft.Security/connectors"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/connectors/aws_dev1"),
		// 			Properties: &armsecurity.ConnectorSettingProperties{
		// 				AuthenticationDetails: &armsecurity.AwsCredsAuthenticationDetailsProperties{
		// 					AuthenticationProvisioningState: to.Ptr(armsecurity.AuthenticationProvisioningStateValid),
		// 					AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsCreds),
		// 					GrantedPermissions: []*armsecurity.PermissionProperty{
		// 						to.Ptr(armsecurity.PermissionPropertyAWSAWSSecurityHubReadOnlyAccess),
		// 						to.Ptr(armsecurity.PermissionPropertyAWSSecurityAudit),
		// 						to.Ptr(armsecurity.PermissionPropertyAWSAmazonSSMAutomationRole)},
		// 						AccountID: to.Ptr("922315681122"),
		// 						AwsAccessKeyID: to.Ptr(""),
		// 						AwsSecretAccessKey: to.Ptr(""),
		// 					},
		// 					HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
		// 						AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
		// 						HybridComputeProvisioningState: to.Ptr(armsecurity.HybridComputeProvisioningStateValid),
		// 						ProxyServer: &armsecurity.ProxyServerProperties{
		// 							IP: to.Ptr("287.221.107.152"),
		// 							Port: to.Ptr("34"),
		// 						},
		// 						Region: to.Ptr("West US 2"),
		// 						ResourceGroupName: to.Ptr("AwsConnectorRG"),
		// 						ServicePrincipal: &armsecurity.ServicePrincipalProperties{
		// 							ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("aws_dev2"),
		// 				Type: to.Ptr("Microsoft.Security/connectors"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/connectors/aws_dev2"),
		// 				Properties: &armsecurity.ConnectorSettingProperties{
		// 					AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
		// 						AuthenticationProvisioningState: to.Ptr(armsecurity.AuthenticationProvisioningStateValid),
		// 						AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsAssumeRole),
		// 						GrantedPermissions: []*armsecurity.PermissionProperty{
		// 							to.Ptr(armsecurity.PermissionPropertyAWSAWSSecurityHubReadOnlyAccess),
		// 							to.Ptr(armsecurity.PermissionPropertyAWSSecurityAudit),
		// 							to.Ptr(armsecurity.PermissionPropertyAWSAmazonSSMAutomationRole)},
		// 							AccountID: to.Ptr("81231569658"),
		// 							AwsAssumeRoleArn: to.Ptr("arn:aws:iam::81231569658:role/AscConnector"),
		// 							AwsExternalID: to.Ptr("20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
		// 						},
		// 						HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
		// 							AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
		// 							HybridComputeProvisioningState: to.Ptr(armsecurity.HybridComputeProvisioningStateValid),
		// 							ProxyServer: &armsecurity.ProxyServerProperties{
		// 								IP: to.Ptr("167.210.187.160"),
		// 								Port: to.Ptr("34"),
		// 							},
		// 							Region: to.Ptr("West US 2"),
		// 							ResourceGroupName: to.Ptr("AwsConnectorRG"),
		// 							ServicePrincipal: &armsecurity.ServicePrincipalProperties{
		// 								ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("gcp_dev"),
		// 					Type: to.Ptr("Microsoft.Security/connectors"),
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/connectors/gcp_dev"),
		// 					Properties: &armsecurity.ConnectorSettingProperties{
		// 						AuthenticationDetails: &armsecurity.GcpCredentialsDetailsProperties{
		// 							AuthenticationProvisioningState: to.Ptr(armsecurity.AuthenticationProvisioningStateValid),
		// 							AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeGcpCredentials),
		// 							GrantedPermissions: []*armsecurity.PermissionProperty{
		// 								to.Ptr(armsecurity.PermissionPropertyGCPSecurityCenterAdminViewer)},
		// 								Type: to.Ptr(""),
		// 								AuthProviderX509CertURL: to.Ptr(""),
		// 								AuthURI: to.Ptr(""),
		// 								ClientEmail: to.Ptr(""),
		// 								ClientID: to.Ptr(""),
		// 								ClientX509CertURL: to.Ptr(""),
		// 								OrganizationID: to.Ptr("AscDemoOrg"),
		// 								PrivateKey: to.Ptr(""),
		// 								PrivateKeyID: to.Ptr(""),
		// 								ProjectID: to.Ptr(""),
		// 								TokenURI: to.Ptr(""),
		// 							},
		// 							HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
		// 								AutoProvision: to.Ptr(armsecurity.AutoProvisionOff),
		// 								HybridComputeProvisioningState: to.Ptr(armsecurity.HybridComputeProvisioningStateInvalid),
		// 								Region: to.Ptr(""),
		// 								ResourceGroupName: to.Ptr(""),
		// 							},
		// 						},
		// 				}},
		// 			}
	}
}
