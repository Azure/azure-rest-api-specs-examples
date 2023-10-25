package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Policy_Create.json
func ExamplePolicyClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPolicyClient().BeginCreate(ctx, "rgrecoveryservicesdatareplication", "4", "fafqwc", armrecoveryservicesdatareplication.PolicyModel{
		Properties: &armrecoveryservicesdatareplication.PolicyModelProperties{
			CustomProperties: &armrecoveryservicesdatareplication.PolicyModelCustomProperties{
				InstanceType: to.Ptr("PolicyModelCustomProperties"),
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
	// res.PolicyModel = armrecoveryservicesdatareplication.PolicyModel{
	// 	Name: to.Ptr("ocmty"),
	// 	Type: to.Ptr("pvltqld"),
	// 	ID: to.Ptr("ffivjzdtqgguqlenedikvdilazliwm"),
	// 	Properties: &armrecoveryservicesdatareplication.PolicyModelProperties{
	// 		CustomProperties: &armrecoveryservicesdatareplication.PolicyModelCustomProperties{
	// 			InstanceType: to.Ptr("PolicyModelCustomProperties"),
	// 		},
	// 		ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armrecoveryservicesdatareplication.PolicyModelSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:54.713Z"); return t}()),
	// 		CreatedBy: to.Ptr("yiaelkrpuzsfumovsxeb"),
	// 		CreatedByType: to.Ptr("qwzrkjsfloruegijrfnfpn"),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:54.713Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("xy"),
	// 		LastModifiedByType: to.Ptr("rnc"),
	// 	},
	// }
}
