package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/CheckfeatureSupport.json
func ExampleClient_CheckFeatureSupport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckFeatureSupport(ctx, "WestUS", &armdataprotection.FeatureValidationRequest{
		ObjectType:  to.Ptr("FeatureValidationRequest"),
		FeatureType: to.Ptr(armdataprotection.FeatureTypeDataSourceType),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.ClientCheckFeatureSupportResponse{
	// 	                            FeatureValidationResponseBaseClassification: &armdataprotection.FeatureValidationResponse{
	// 		ObjectType: to.Ptr("FeatureValidationResponse"),
	// 		FeatureType: to.Ptr(armdataprotection.FeatureTypeDataSourceType),
	// 		Features: []*armdataprotection.SupportedFeature{
	// 			{
	// 				ExposureControlledFeatures: []*string{
	// 				},
	// 				FeatureName: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	// 				SupportStatus: to.Ptr(armdataprotection.FeatureSupportStatusPrivatePreview),
	// 			},
	// 			{
	// 				ExposureControlledFeatures: []*string{
	// 				},
	// 				FeatureName: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 				SupportStatus: to.Ptr(armdataprotection.FeatureSupportStatusPublicPreview),
	// 		}},
	// 	},
	// 	                        }
}
