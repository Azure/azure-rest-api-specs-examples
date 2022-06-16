package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobMitigate.json
func ExampleManagementClient_Mitigate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewManagementClient("<subscription-id>", cred, nil)
	_, err = client.Mitigate(ctx,
		"<job-name>",
		"<resource-group-name>",
		armdatabox.MitigateJobRequest{
			CustomerResolutionCode: armdatabox.CustomerResolutionCodeMoveToCleanUpDevice.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
