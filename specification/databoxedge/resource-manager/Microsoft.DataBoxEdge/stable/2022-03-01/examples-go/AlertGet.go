package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/AlertGet.json
func ExampleAlertsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().Get(ctx, "testedgedevice", "159a00c7-8543-4343-9435-263ac87df3bb", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Alert = armdataboxedge.Alert{
	// 	Name: to.Ptr("159a00c7-8543-4343-9435-263ac87df3bb"),
	// 	Type: to.Ptr("dataBoxEdgeDevices/alerts"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/alerts/159a00c7-8543-4343-9435-263ac87df3bb"),
	// 	Properties: &armdataboxedge.AlertProperties{
	// 		AlertType: to.Ptr("UpdateScanFailedEvent"),
	// 		AppearedAtDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:18:51.427Z"); return t}()),
	// 		DetailedInformation: map[string]*string{
	// 		},
	// 		ErrorDetails: &armdataboxedge.AlertErrorDetails{
	// 			ErrorCode: to.Ptr(""),
	// 			ErrorMessage: to.Ptr(""),
	// 			Occurrences: to.Ptr[int32](1),
	// 		},
	// 		Recommendation: to.Ptr("Resolve the error : An internal error has occurred. Please contact Microsoft Support."),
	// 		Severity: to.Ptr(armdataboxedge.AlertSeverityCritical),
	// 		Title: to.Ptr("Could not scan for updates. Error message : 'An internal error has occurred. Please contact Microsoft Support.'."),
	// 	},
	// }
}
