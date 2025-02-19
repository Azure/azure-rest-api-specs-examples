package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetUpdates.json
func ExampleUpdatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpdatesClient().Get(ctx, "testrg", "testcluster", "Microsoft4.2203.2.32", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Update = armazurestackhci.Update{
	// 	Name: to.Ptr("Microsoft4.2203.2.32"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/updates"),
	// 	ID: to.Ptr("/subscriptions/b8d594e5-51f3-4c11-9c54-a7771b81c712/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/clusters/testcluster/updates/Microsoft4.2203.2.32"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armazurestackhci.UpdateProperties{
	// 		Description: to.Ptr("AzS Update 4.2203.2.32"),
	// 		AdditionalProperties: to.Ptr("additional properties"),
	// 		AvailabilityType: to.Ptr(armazurestackhci.AvailabilityTypeLocal),
	// 		DisplayName: to.Ptr("AzS Update - 4.2203.2.32"),
	// 		InstalledDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T14:08:18.254Z"); return t}()),
	// 		PackagePath: to.Ptr("\\\\SU1FileServer\\SU1_Infrastructure_2\\Updates\\Packages\\Microsoft4.2203.2.32"),
	// 		PackageSizeInMb: to.Ptr[float32](18858),
	// 		PackageType: to.Ptr("Infrastructure"),
	// 		Prerequisites: []*armazurestackhci.UpdatePrerequisite{
	// 			{
	// 				PackageName: to.Ptr("update package name"),
	// 				UpdateType: to.Ptr("update type"),
	// 				Version: to.Ptr("prerequisite version"),
	// 		}},
	// 		Publisher: to.Ptr("Microsoft"),
	// 		ReleaseLink: to.Ptr("https://docs.microsoft.com/azure-stack/operator/release-notes?view=azs-2203"),
	// 		State: to.Ptr(armazurestackhci.StateInstalled),
	// 		UpdateStateProperties: &armazurestackhci.UpdateStateProperties{
	// 			NotifyMessage: to.Ptr("Brief message with instructions for updates of AvailabilityType Notify"),
	// 			ProgressPercentage: to.Ptr[float32](0),
	// 		},
	// 		Version: to.Ptr("4.2203.2.32"),
	// 	},
	// }
}
