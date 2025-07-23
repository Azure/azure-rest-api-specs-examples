package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/AuthorizedApplications_Get.json
func ExampleAuthorizedApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizedApplicationsClient().Get(ctx, "Microsoft.Contoso", "760505bf-dcfa-4311-b890-18da392a00b2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizedApplication = armproviderhub.AuthorizedApplication{
	// 	Name: to.Ptr("Microsoft.Contoso/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/authorizedApplications"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/authorizedApplications/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.AuthorizedApplicationProperties{
	// 		DataAuthorizations: []*armproviderhub.ApplicationDataAuthorization{
	// 			{
	// 				ResourceTypes: []*string{
	// 					to.Ptr("*")},
	// 					Role: to.Ptr(armproviderhub.RoleServiceOwner),
	// 			}},
	// 			ProviderAuthorization: &armproviderhub.ApplicationProviderAuthorization{
	// 				ManagedByRoleDefinitionID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 				RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 			},
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
