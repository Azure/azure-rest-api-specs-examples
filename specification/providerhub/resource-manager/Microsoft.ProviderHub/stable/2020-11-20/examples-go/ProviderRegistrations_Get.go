package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_Get.json
func ExampleProviderRegistrationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderRegistrationsClient().Get(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderRegistration = armproviderhub.ProviderRegistration{
	// 	Properties: &armproviderhub.ProviderRegistrationProperties{
	// 		Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 			{
	// 				Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 				QuotaID: to.Ptr("CSP_2015-05-01"),
	// 			},
	// 			{
	// 				Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 				QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 		}},
	// 		Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
	// 			IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 			IncidentRoutingService: to.Ptr(""),
	// 			IncidentRoutingTeam: to.Ptr(""),
	// 			ManifestOwners: []*string{
	// 				to.Ptr("SPARTA-PlatformServiceAdministrator")},
	// 				ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceProviderManagementResourceAccessPolicyNotSpecified),
	// 			},
	// 			Namespace: to.Ptr("microsoft.contoso"),
	// 			ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 				{
	// 					ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 					RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 			}},
	// 			ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 			ProviderVersion: to.Ptr("2.0"),
	// 			ProviderHubMetadata: &armproviderhub.ProviderRegistrationPropertiesProviderHubMetadata{
	// 				ProviderAuthentication: &armproviderhub.MetadataProviderAuthentication{
	// 					AllowedAudiences: []*string{
	// 						to.Ptr("https://management.core.windows.net/")},
	// 					},
	// 				},
	// 				ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 			},
	// 		}
}
