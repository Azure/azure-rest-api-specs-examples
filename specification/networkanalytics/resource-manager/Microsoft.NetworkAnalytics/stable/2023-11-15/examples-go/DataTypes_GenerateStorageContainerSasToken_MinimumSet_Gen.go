package armnetworkanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataTypes_GenerateStorageContainerSasToken_MinimumSet_Gen.json
func ExampleDataTypesClient_GenerateStorageContainerSasToken_dataTypesGenerateStorageContainerSasTokenMaximumSetGenGeneratedByMinimumSetRuleMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataTypesClient().GenerateStorageContainerSasToken(ctx, "aoiresourceGroupName", "dataproduct01", "datatypename", armnetworkanalytics.ContainerSaS{
		ExpiryTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T05:35:16.887Z"); return t }()),
		IPAddress:       to.Ptr("1.1.1.1"),
		StartTimeStamp:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T05:35:16.887Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerSasToken = armnetworkanalytics.ContainerSasToken{
	// 	StorageContainerSasToken: to.Ptr("storageContainerSasToken"),
	// }
}
