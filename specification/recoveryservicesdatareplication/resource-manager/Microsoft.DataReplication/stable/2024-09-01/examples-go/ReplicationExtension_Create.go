package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/ReplicationExtension_Create.json
func ExampleReplicationExtensionClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationExtensionClient().BeginCreate(ctx, "rgrecoveryservicesdatareplication", "4", "g16yjJ", armrecoveryservicesdatareplication.ReplicationExtensionModel{
		Properties: &armrecoveryservicesdatareplication.ReplicationExtensionModelProperties{
			CustomProperties: &armrecoveryservicesdatareplication.HyperVToAzStackHCIReplicationExtensionModelCustomProperties{
				InstanceType: to.Ptr("ReplicationExtensionModelCustomProperties"),
			},
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
	// res = armrecoveryservicesdatareplication.ReplicationExtensionClientCreateResponse{
	// 	ReplicationExtensionModel: &armrecoveryservicesdatareplication.ReplicationExtensionModel{
	// 		Properties: &armrecoveryservicesdatareplication.ReplicationExtensionModelProperties{
	// 			ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
	// 			CustomProperties: &armrecoveryservicesdatareplication.HyperVToAzStackHCIReplicationExtensionModelCustomProperties{
	// 				InstanceType: to.Ptr("ReplicationExtensionModelCustomProperties"),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/replicationExtensions/extension1"),
	// 		Name: to.Ptr("xvjffbiecsd"),
	// 		Type: to.Ptr("miadbgilpheilzfoonveznybthgdwh"),
	// 		SystemData: &armrecoveryservicesdatareplication.SystemData{
	// 			CreatedBy: to.Ptr("lagtinqhksctfdxmfbpf"),
	// 			CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("dsqllpglanwztdmisrknjtqz")),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("suwjpejlaya"),
	// 			LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("nrfjuyghtbivwihr")),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
	// 		},
	// 	},
	// }
}
