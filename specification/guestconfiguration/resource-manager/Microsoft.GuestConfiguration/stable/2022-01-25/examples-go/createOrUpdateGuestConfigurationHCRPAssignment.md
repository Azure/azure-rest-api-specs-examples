```go
package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/createOrUpdateGuestConfigurationHCRPAssignment.json
func ExampleHCRPAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armguestconfiguration.NewHCRPAssignmentsClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"NotInstalledApplicationForWindows",
		"myResourceGroupName",
		"myMachineName",
		armguestconfiguration.Assignment{
			Name:     to.Ptr("NotInstalledApplicationForWindows"),
			Location: to.Ptr("westcentralus"),
			Properties: &armguestconfiguration.AssignmentProperties{
				Context: to.Ptr("Azure policy"),
				GuestConfiguration: &armguestconfiguration.Navigation{
					Name:           to.Ptr("NotInstalledApplicationForWindows"),
					AssignmentType: to.Ptr(armguestconfiguration.AssignmentTypeApplyAndAutoCorrect),
					ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
						{
							Name:  to.Ptr("[InstalledApplication]NotInstalledApplicationResource1;Name"),
							Value: to.Ptr("NotePad,sql"),
						}},
					ContentHash: to.Ptr("123contenthash"),
					ContentURI:  to.Ptr("https://thisisfake/pacakge"),
					Version:     to.Ptr("1.*"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fguestconfiguration%2Farmguestconfiguration%2Fv1.0.0/sdk/resourcemanager/guestconfiguration/armguestconfiguration/README.md) on how to add the SDK to your project and authenticate.
