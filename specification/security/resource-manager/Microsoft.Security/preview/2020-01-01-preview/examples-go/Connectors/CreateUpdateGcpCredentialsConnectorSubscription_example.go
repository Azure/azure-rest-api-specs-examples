package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateGcpCredentialsConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate_gcpCredentialsCreateACloudAccountConnectorForASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountConnectorsClient().CreateOrUpdate(ctx, "gcp_dev", armsecurity.ConnectorSetting{
		Properties: &armsecurity.ConnectorSettingProperties{
			AuthenticationDetails: &armsecurity.GcpCredentialsDetailsProperties{
				AuthenticationType:      to.Ptr(armsecurity.AuthenticationTypeGcpCredentials),
				Type:                    to.Ptr("service_account"),
				AuthProviderX509CertURL: to.Ptr("https://www.googleapis.com/oauth2/v1/certs"),
				AuthURI:                 to.Ptr("https://accounts.google.com/o/oauth2/auth"),
				ClientEmail:             to.Ptr("asc-135@asc-project-1234.iam.gserviceaccount.com"),
				ClientID:                to.Ptr("105889053725632919854"),
				ClientX509CertURL:       to.Ptr("https://www.googleapis.com/robot/v1/metadata/x509/asc-135%40asc-project-1234.iam.gserviceaccount.com"),
				OrganizationID:          to.Ptr("AscDemoOrg"),
				PrivateKey:              to.Ptr("******"),
				PrivateKeyID:            to.Ptr("6efg587hra2568as34d22326b044cc20dc2af"),
				ProjectID:               to.Ptr("asc-project-1234"),
				TokenURI:                to.Ptr("https://oauth2.googleapis.com/token"),
			},
			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
				AutoProvision: to.Ptr(armsecurity.AutoProvisionOff),
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
	// 	Name: to.Ptr("gcp_dev"),
	// 	Type: to.Ptr("Microsoft.Security/connectors"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/connectors/gcp_dev"),
	// 	Properties: &armsecurity.ConnectorSettingProperties{
	// 		AuthenticationDetails: &armsecurity.GcpCredentialsDetailsProperties{
	// 			AuthenticationProvisioningState: to.Ptr(armsecurity.AuthenticationProvisioningStateValid),
	// 			AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeGcpCredentials),
	// 			GrantedPermissions: []*armsecurity.PermissionProperty{
	// 				to.Ptr(armsecurity.PermissionPropertyGCPSecurityCenterAdminViewer)},
	// 				Type: to.Ptr(""),
	// 				AuthProviderX509CertURL: to.Ptr(""),
	// 				AuthURI: to.Ptr(""),
	// 				ClientEmail: to.Ptr(""),
	// 				ClientID: to.Ptr(""),
	// 				ClientX509CertURL: to.Ptr(""),
	// 				OrganizationID: to.Ptr("AscDemoOrg"),
	// 				PrivateKey: to.Ptr(""),
	// 				PrivateKeyID: to.Ptr(""),
	// 				ProjectID: to.Ptr(""),
	// 				TokenURI: to.Ptr(""),
	// 			},
	// 			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
	// 				AutoProvision: to.Ptr(armsecurity.AutoProvisionOff),
	// 				HybridComputeProvisioningState: to.Ptr(armsecurity.HybridComputeProvisioningStateInvalid),
	// 				Region: to.Ptr(""),
	// 				ResourceGroupName: to.Ptr(""),
	// 			},
	// 		},
	// 	}
}
