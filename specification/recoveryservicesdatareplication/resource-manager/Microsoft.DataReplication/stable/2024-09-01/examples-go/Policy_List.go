package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/Policy_List.json
func ExamplePolicyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyClient().NewListPager("rgrecoveryservicesdatareplication", "4", nil)
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
		// page = armrecoveryservicesdatareplication.PolicyClientListResponse{
		// 	PolicyModelListResult: armrecoveryservicesdatareplication.PolicyModelListResult{
		// 		Value: []*armrecoveryservicesdatareplication.PolicyModel{
		// 			{
		// 				Properties: &armrecoveryservicesdatareplication.PolicyModelProperties{
		// 					ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
		// 					CustomProperties: &armrecoveryservicesdatareplication.HyperVToAzStackHCIPolicyModelCustomProperties{
		// 						InstanceType: to.Ptr("PolicyModelCustomProperties"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/replicationPolicies/policy1"),
		// 				Name: to.Ptr("ocmty"),
		// 				Type: to.Ptr("pvltqld"),
		// 				SystemData: &armrecoveryservicesdatareplication.SystemData{
		// 					CreatedBy: to.Ptr("yiaelkrpuzsfumovsxeb"),
		// 					CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("qwzrkjsfloruegijrfnfpn")),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:54.713Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("xy"),
		// 					LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("rnc")),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:54.713Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
