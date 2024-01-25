package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_ListByStreamingJob_Diagnostics.json
func ExampleInputsClient_NewListByStreamingJobPager_listAllInputsInAStreamingJobAndIncludeDiagnosticInformationUsingTheSelectODataQueryParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInputsClient().NewListByStreamingJobPager("sjrg3276", "sj7804", &armstreamanalytics.InputsClientListByStreamingJobOptions{Select: to.Ptr("*")})
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
		// page.InputListResult = armstreamanalytics.InputListResult{
		// 	Value: []*armstreamanalytics.Input{
		// 		{
		// 			Name: to.Ptr("inputtest"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/inputs/inputtest"),
		// 			Properties: &armstreamanalytics.StreamInputProperties{
		// 				Type: to.Ptr("Stream"),
		// 				Diagnostics: &armstreamanalytics.Diagnostics{
		// 					Conditions: []*armstreamanalytics.DiagnosticCondition{
		// 						{
		// 							Code: to.Ptr("INP-3"),
		// 							Message: to.Ptr("Could not deserialize the input event as Json. Some possible reasons: 1) Malformed events 2) Input source configured with incorrect serialization format"),
		// 							Since: to.Ptr("2017-05-11T04:38:42.4938687Z"),
		// 					}},
		// 				},
		// 				Etag: to.Ptr("ca88f8fa-605b-4c7f-8695-46f5faa60cd0"),
		// 				Serialization: &armstreamanalytics.JSONSerialization{
		// 					Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
		// 					Properties: &armstreamanalytics.JSONSerializationProperties{
		// 						Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 					},
		// 				},
		// 				Datasource: &armstreamanalytics.BlobStreamInputDataSource{
		// 					Type: to.Ptr("Microsoft.Storage/Blob"),
		// 					Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
		// 						Container: to.Ptr("state"),
		// 						PathPattern: to.Ptr(""),
		// 						StorageAccounts: []*armstreamanalytics.StorageAccount{
		// 							{
		// 								AccountName: to.Ptr("someAccountName"),
		// 						}},
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
