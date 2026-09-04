package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/AuthorizedApplications_Get.json
func ExampleAuthorizedApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
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
	// res = armproviderhub.AuthorizedApplicationsClientGetResponse{
	// 	AuthorizedApplication: armproviderhub.AuthorizedApplication{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/authorizedApplications/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 		Name: to.Ptr("Microsoft.Contoso/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/authorizedApplications"),
	// 		Properties: &armproviderhub.AuthorizedApplicationProperties{
	// 			ProviderAuthorization: &armproviderhub.ApplicationProviderAuthorization{
	// 				RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 				ManagedByRoleDefinitionID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 			},
	// 			DataAuthorizations: []*armproviderhub.ApplicationDataAuthorization{
	// 				{
	// 					Role: to.Ptr(armproviderhub.RoleServiceOwner),
	// 					ResourceTypes: []*string{
	// 						to.Ptr("*"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armproviderhub.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 		},
	// 	},
	// }
}
