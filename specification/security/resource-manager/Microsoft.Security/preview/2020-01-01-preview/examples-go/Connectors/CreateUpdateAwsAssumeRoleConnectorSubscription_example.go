package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate_awsAssumeRoleCreateACloudAccountConnectorForASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountConnectorsClient().CreateOrUpdate(ctx, "aws_dev2", armsecurity.ConnectorSetting{
		Properties: &armsecurity.ConnectorSettingProperties{
			AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
				AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsAssumeRole),
				AwsAssumeRoleArn:   to.Ptr("arn:aws:iam::81231569658:role/AscConnector"),
				AwsExternalID:      to.Ptr("20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectorSetting = armsecurity.ConnectorSetting{
	// 	Name: to.Ptr("aws_dev2"),
	// 	Type: to.Ptr("Microsoft.Security/connectors"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/connectors/aws_dev2"),
	// 	Properties: &armsecurity.ConnectorSettingProperties{
	// 		AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
	// 			AuthenticationProvisioningState: to.Ptr(armsecurity.AuthenticationProvisioningStateValid),
	// 			AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsAssumeRole),
	// 			GrantedPermissions: []*armsecurity.PermissionProperty{
	// 				to.Ptr(armsecurity.PermissionPropertyAWSAWSSecurityHubReadOnlyAccess),
	// 				to.Ptr(armsecurity.PermissionPropertyAWSSecurityAudit),
	// 				to.Ptr(armsecurity.PermissionPropertyAWSAmazonSSMAutomationRole)},
	// 				AccountID: to.Ptr("81231569658"),
	// 				AwsAssumeRoleArn: to.Ptr("arn:aws:iam::81231569658:role/AscConnector"),
	// 				AwsExternalID: to.Ptr("20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
	// 			},
	// 			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
	// 				AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
	// 				HybridComputeProvisioningState: to.Ptr(armsecurity.HybridComputeProvisioningStateValid),
	// 				ProxyServer: &armsecurity.ProxyServerProperties{
	// 					IP: to.Ptr("167.220.197.140"),
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
