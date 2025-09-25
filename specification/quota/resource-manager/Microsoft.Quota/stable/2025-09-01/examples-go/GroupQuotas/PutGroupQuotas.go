package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota/v2"
)

// Generated from example definition: 2025-09-01/GroupQuotas/PutGroupQuotas.json
func ExampleGroupQuotasClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGroupQuotasClient().BeginCreateOrUpdate(ctx, "E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", armquota.GroupQuotasEntity{
		Properties: &armquota.GroupQuotasEntityProperties{
			DisplayName: to.Ptr("GroupQuota1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armquota.GroupQuotasClientCreateOrUpdateResponse{
	// 	GroupQuotasEntity: &armquota.GroupQuotasEntity{
	// 		Name: to.Ptr("groupquota1"),
	// 		Type: to.Ptr("Microsoft.Quota/groupQuotas"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1"),
	// 		Properties: &armquota.GroupQuotasEntityProperties{
	// 			DisplayName: to.Ptr("GroupQuota1"),
	// 			GroupType: to.Ptr(armquota.GroupTypeAllocationGroup),
	// 			ProvisioningState: to.Ptr(armquota.RequestStateAccepted),
	// 		},
	// 	},
	// }
}
