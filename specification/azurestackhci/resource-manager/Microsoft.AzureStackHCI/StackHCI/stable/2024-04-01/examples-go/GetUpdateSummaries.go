package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetUpdateSummaries.json
func ExampleUpdateSummariesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpdateSummariesClient().Get(ctx, "testrg", "testcluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpdateSummaries = armazurestackhci.UpdateSummaries{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/updateSummaries"),
	// 	ID: to.Ptr("/subscriptions/b8d594e5-51f3-4c11-9c54-a7771b81c712/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/clusters/testcluster/updateSummaries/default"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armazurestackhci.UpdateSummariesProperties{
	// 		CurrentVersion: to.Ptr("4.2203.2.32"),
	// 		HardwareModel: to.Ptr("PowerEdge R730xd"),
	// 		OemFamily: to.Ptr("DellEMC"),
	// 		PackageVersions: []*armazurestackhci.PackageVersionInfo{
	// 			{
	// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-07T18:04:07.000Z"); return t}()),
	// 				PackageType: to.Ptr("OEM"),
	// 				Version: to.Ptr("2.2.2108.6"),
	// 			},
	// 			{
	// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-07T18:04:07.000Z"); return t}()),
	// 				PackageType: to.Ptr("Services"),
	// 				Version: to.Ptr("4.2203.2.32"),
	// 			},
	// 			{
	// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-07T18:04:07.000Z"); return t}()),
	// 				PackageType: to.Ptr("Infrastructure"),
	// 				Version: to.Ptr("4.2203.2.32"),
	// 		}},
	// 		State: to.Ptr(armazurestackhci.UpdateSummariesPropertiesStateAppliedSuccessfully),
	// 	},
	// }
}
