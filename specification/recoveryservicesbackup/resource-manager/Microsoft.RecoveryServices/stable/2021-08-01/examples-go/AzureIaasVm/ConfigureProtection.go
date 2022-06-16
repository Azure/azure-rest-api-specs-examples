package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureIaasVm/ConfigureProtection.json
func ExampleProtectedItemsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectedItemsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		"<protected-item-name>",
		armrecoveryservicesbackup.ProtectedItemResource{
			Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
				AzureIaaSVMProtectedItem: armrecoveryservicesbackup.AzureIaaSVMProtectedItem{
					ProtectedItem: armrecoveryservicesbackup.ProtectedItem{
						PolicyID:          to.StringPtr("<policy-id>"),
						ProtectedItemType: to.StringPtr("<protected-item-type>"),
						SourceResourceID:  to.StringPtr("<source-resource-id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProtectedItemResource.ID: %s\n", *res.ID)
}
