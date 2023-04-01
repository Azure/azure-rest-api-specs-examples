package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_AzureDataLakeStore_JSON.json
func ExampleOutputsClient_CreateOrReplace_createAnAzureDataLakeStoreOutputWithJsonSerialization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().CreateOrReplace(ctx, "sjrg6912", "sj3310", "output5195", armstreamanalytics.Output{
		Properties: &armstreamanalytics.OutputProperties{
			Datasource: &armstreamanalytics.AzureDataLakeStoreOutputDataSource{
				Type: to.Ptr("Microsoft.DataLake/Accounts"),
				Properties: &armstreamanalytics.AzureDataLakeStoreOutputDataSourceProperties{
					RefreshToken:           to.Ptr("someRefreshToken=="),
					TokenUserDisplayName:   to.Ptr("Bob Smith"),
					TokenUserPrincipalName: to.Ptr("bobsmith@contoso.com"),
					AccountName:            to.Ptr("someaccount"),
					DateFormat:             to.Ptr("yyyy/MM/dd"),
					FilePathPrefix:         to.Ptr("{date}/{time}"),
					TenantID:               to.Ptr("cea4e98b-c798-49e7-8c40-4a2b3beb47dd"),
					TimeFormat:             to.Ptr("HH"),
				},
			},
			Serialization: &armstreamanalytics.JSONSerialization{
				Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
				Properties: &armstreamanalytics.JSONSerializationProperties{
					Format:   to.Ptr(armstreamanalytics.JSONOutputSerializationFormatArray),
					Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
				},
			},
		},
	}, &armstreamanalytics.OutputsClientCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("output5195"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg6912/providers/Microsoft.StreamAnalytics/streamingjobs/sj3310/outputs/output5195"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.AzureDataLakeStoreOutputDataSource{
	// 			Type: to.Ptr("Microsoft.DataLake/Accounts"),
	// 			Properties: &armstreamanalytics.AzureDataLakeStoreOutputDataSourceProperties{
	// 				TokenUserDisplayName: to.Ptr("Bob Smith"),
	// 				TokenUserPrincipalName: to.Ptr("bobsmith@contoso.com"),
	// 				AccountName: to.Ptr("someaccount"),
	// 				DateFormat: to.Ptr("yyyy/MM/dd"),
	// 				FilePathPrefix: to.Ptr("{date}/{time}"),
	// 				TenantID: to.Ptr("cea4e98b-c798-49e7-8c40-4a2b3beb47dd"),
	// 				TimeFormat: to.Ptr("HH"),
	// 			},
	// 		},
	// 		Serialization: &armstreamanalytics.JSONSerialization{
	// 			Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
	// 			Properties: &armstreamanalytics.JSONSerializationProperties{
	// 				Format: to.Ptr(armstreamanalytics.JSONOutputSerializationFormatArray),
	// 				Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
	// 			},
	// 		},
	// 	},
	// }
}
