package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/AlertGetAllInDevice.json
func ExampleAlertsClient_NewListByDataBoxEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsClient().NewListByDataBoxEdgeDevicePager("testedgedevice", "GroupForEdgeAutomation", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AlertList = armdataboxedge.AlertList{
		// 	Value: []*armdataboxedge.Alert{
		// 		{
		// 			Name: to.Ptr("83eccd0b-134b-40b0-ad62-b5f124d03790"),
		// 			Type: to.Ptr("dataBoxEdgeDevices/alerts"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/alerts/83eccd0b-134b-40b0-ad62-b5f124d03790"),
		// 			Properties: &armdataboxedge.AlertProperties{
		// 				AlertType: to.Ptr("PasswordChangedEvent"),
		// 				AppearedAtDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:29:15.232Z"); return t}()),
		// 				DetailedInformation: map[string]*string{
		// 				},
		// 				ErrorDetails: &armdataboxedge.AlertErrorDetails{
		// 					ErrorCode: to.Ptr(""),
		// 					ErrorMessage: to.Ptr(""),
		// 					Occurrences: to.Ptr[int32](2),
		// 				},
		// 				Recommendation: to.Ptr("The device administrator password has changed. This is a required action as part of the first time device setup or regular password reset. No further action is required."),
		// 				Severity: to.Ptr(armdataboxedge.AlertSeverityInformational),
		// 				Title: to.Ptr("Device password has changed"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("159a00c7-8543-4343-9435-263ac87df3bb"),
		// 			Type: to.Ptr("dataBoxEdgeDevices/alerts"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/alerts/159a00c7-8543-4343-9435-263ac87df3bb"),
		// 			Properties: &armdataboxedge.AlertProperties{
		// 				AlertType: to.Ptr("UpdateScanFailedEvent"),
		// 				AppearedAtDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:18:51.427Z"); return t}()),
		// 				DetailedInformation: map[string]*string{
		// 				},
		// 				ErrorDetails: &armdataboxedge.AlertErrorDetails{
		// 					ErrorCode: to.Ptr(""),
		// 					ErrorMessage: to.Ptr(""),
		// 					Occurrences: to.Ptr[int32](1),
		// 				},
		// 				Recommendation: to.Ptr("Resolve the error : An internal error has occurred. Please contact Microsoft Support."),
		// 				Severity: to.Ptr(armdataboxedge.AlertSeverityCritical),
		// 				Title: to.Ptr("Could not scan for updates. Error message : 'An internal error has occurred. Please contact Microsoft Support.'."),
		// 			},
		// 	}},
		// }
	}
}
