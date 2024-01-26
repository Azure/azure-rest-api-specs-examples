package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Get_PowerBI.json
func ExampleOutputsClient_Get_getAPowerBiOutput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().Get(ctx, "sjrg7983", "sj2331", "output3022", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("output3022"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg7983/providers/Microsoft.StreamAnalytics/streamingjobs/sj2331/outputs/output3022"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.PowerBIOutputDataSource{
	// 			Type: to.Ptr("PowerBI"),
	// 			Properties: &armstreamanalytics.PowerBIOutputDataSourceProperties{
	// 				TokenUserDisplayName: to.Ptr("Bob Smith"),
	// 				TokenUserPrincipalName: to.Ptr("bobsmith@contoso.com"),
	// 				Dataset: to.Ptr("someDataset"),
	// 				GroupID: to.Ptr("ac40305e-3e8d-43ac-8161-c33799f43e95"),
	// 				GroupName: to.Ptr("MyPowerBIGroup"),
	// 				Table: to.Ptr("someTable"),
	// 			},
	// 		},
	// 	},
	// }
}
