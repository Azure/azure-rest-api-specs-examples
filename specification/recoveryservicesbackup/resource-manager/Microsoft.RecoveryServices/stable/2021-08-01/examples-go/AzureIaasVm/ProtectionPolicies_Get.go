package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureIaasVm/ProtectionPolicies_Get.json
func ExampleProtectionPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectionPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProtectionPolicyResource.ID: %s\n", *res.ID)
}
