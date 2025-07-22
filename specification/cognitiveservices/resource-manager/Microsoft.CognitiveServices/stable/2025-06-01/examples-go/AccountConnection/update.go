package armcognitiveservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/AccountConnection/update.json
func ExampleAccountConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountConnectionsClient().Update(ctx, "test-rg", "account-1", "connection-1", &armcognitiveservices.AccountConnectionsClientUpdateOptions{Connection: &armcognitiveservices.ConnectionUpdateContent{
		Properties: &armcognitiveservices.AccessKeyAuthTypeConnectionProperties{
			AuthType:   to.Ptr(armcognitiveservices.ConnectionAuthTypeAccessKey),
			Category:   to.Ptr(armcognitiveservices.ConnectionCategoryADLSGen2),
			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00.000Z"); return t }()),
			Metadata:   map[string]*string{},
			Target:     to.Ptr("some_string"),
			Credentials: &armcognitiveservices.ConnectionAccessKey{
				AccessKeyID:     to.Ptr("some_string"),
				SecretAccessKey: to.Ptr("some_string"),
			},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionPropertiesV2BasicResource = armcognitiveservices.ConnectionPropertiesV2BasicResource{
	// 	Name: to.Ptr("connection-1"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/connections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.CognitiveServices/accounts/account-1/connections/connection-1"),
	// 	Properties: &armcognitiveservices.AccessKeyAuthTypeConnectionProperties{
	// 		AuthType: to.Ptr(armcognitiveservices.ConnectionAuthTypeAccessKey),
	// 		Category: to.Ptr(armcognitiveservices.ConnectionCategoryADLSGen2),
	// 		ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00.000Z"); return t}()),
	// 		Metadata: map[string]*string{
	// 		},
	// 		Target: to.Ptr("some_string"),
	// 	},
	// }
}
