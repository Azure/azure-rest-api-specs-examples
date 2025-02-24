package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/BmcKeySets_ListByCluster.json
func ExampleBmcKeySetsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBmcKeySetsClient().NewListByClusterPager("resourceGroupName", "clusterName", nil)
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
		// page.BmcKeySetList = armnetworkcloud.BmcKeySetList{
		// 	Value: []*armnetworkcloud.BmcKeySet{
		// 		{
		// 			Name: to.Ptr("bmcKeySetName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/clusters/bmcKeySets"),
		// 			ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bmcKeySets/bmcKeySetName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.BmcKeySetProperties{
		// 				AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
		// 				DetailedStatus: to.Ptr(armnetworkcloud.BmcKeySetDetailedStatusSomeInvalid),
		// 				DetailedStatusMessage: to.Ptr("Invalid Azure user(s) were provided: userXYZ"),
		// 				Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
		// 				LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
		// 				PrivilegeLevel: to.Ptr(armnetworkcloud.BmcKeySetPrivilegeLevelAdministrator),
		// 				ProvisioningState: to.Ptr(armnetworkcloud.BmcKeySetProvisioningStateSucceeded),
		// 				UserList: []*armnetworkcloud.KeySetUser{
		// 					{
		// 						Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
		// 						AzureUserName: to.Ptr("userABC"),
		// 						SSHPublicKey: &armnetworkcloud.SSHPublicKey{
		// 							KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 						},
		// 						UserPrincipalName: to.Ptr("userABC@contoso.com"),
		// 					},
		// 					{
		// 						Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
		// 						AzureUserName: to.Ptr("userXYZ"),
		// 						SSHPublicKey: &armnetworkcloud.SSHPublicKey{
		// 							KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 						},
		// 						UserPrincipalName: to.Ptr("userXYZ@contoso.com"),
		// 				}},
		// 				UserListStatus: []*armnetworkcloud.KeySetUserStatus{
		// 					{
		// 						AzureUserName: to.Ptr("userABC"),
		// 						Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
		// 						StatusMessage: to.Ptr("User has been provisioned"),
		// 					},
		// 					{
		// 						AzureUserName: to.Ptr("userXYZ"),
		// 						Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
		// 						StatusMessage: to.Ptr("User is not a valid Azure user"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
