package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53d56e4ec74156c450d1e51745a971d3f2031dd7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/WorkspacesUpdate.json
func ExampleWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Update(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.WorkspacePatch{
		Properties: &armoperationalinsights.WorkspaceProperties{
			RetentionInDays: to.Ptr[int32](30),
			SKU: &armoperationalinsights.WorkspaceSKU{
				Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
			},
			WorkspaceCapping: &armoperationalinsights.WorkspaceCapping{
				DailyQuotaGb: to.Ptr[float64](-1),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armoperationalinsights.Workspace{
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("val1"),
	// 	},
	// 	Properties: &armoperationalinsights.WorkspaceProperties{
	// 		PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		RetentionInDays: to.Ptr[int32](30),
	// 		SKU: &armoperationalinsights.WorkspaceSKU{
	// 			Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
	// 		},
	// 		WorkspaceCapping: &armoperationalinsights.WorkspaceCapping{
	// 			DailyQuotaGb: to.Ptr[float64](-1),
	// 			DataIngestionStatus: to.Ptr(armoperationalinsights.DataIngestionStatusRespectQuota),
	// 			QuotaNextResetTime: to.Ptr("Mon, 16 Nov 2020 15:00:00 GMT"),
	// 		},
	// 	},
	// }
}
