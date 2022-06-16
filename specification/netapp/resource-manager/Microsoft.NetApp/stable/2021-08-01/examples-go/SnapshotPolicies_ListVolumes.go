package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_ListVolumes.json
func ExampleSnapshotPoliciesClient_ListVolumes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.ListVolumes(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientListVolumesResult)
}
