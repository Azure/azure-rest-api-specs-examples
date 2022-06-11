```go
package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/StorageAccountCredentialsCreateOrUpdate.json
func ExampleStorageAccountCredentialsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewStorageAccountCredentialsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"DummySacForSDKTest",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.StorageAccountCredential{
			Name: to.Ptr("DummySacForSDKTest"),
			Properties: &armstorsimple1200series.StorageAccountCredentialProperties{
				AccessKey: &armstorsimple1200series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:             to.Ptr(armstorsimple1200series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertificateThumbprint: to.Ptr("D73DB57C4CDD6761E159F8D1E8A7D759424983FD"),
					Value:                           to.Ptr("Ev1tm0QBmpGGm4a58GkqLqx8veJEEgQtg5K3Jizpmy7JdSv9dlcRwk59THw6KIdMDlEHcS8mPyneBtOEQsh4wkcFB7qrmQz+KsRAyIhEm6bwPEm3qN8+aDDzNcXn/6vu/sqV0AP7zit9/s7SxXGxjKrz4zKnOy16/DbzRRmUHNO+HO6JUM0cUfHXTX0mEecbsXqBq0A8IEG8z+bJgXX1EhoGkzE6yVsObm4S1AcKrLiwWjqmSLji5Q8gGO+y4KTTmC3p45h5GHHXjJyOccHhySWDAffxnTzUD/sOoh+aD2VkAYrL3DdnkVzhAdfcZfVI4soONx7tYMloZIVsfW1M2Q=="),
				},
				CloudType: to.Ptr(armstorsimple1200series.CloudTypeAzure),
				EnableSSL: to.Ptr(armstorsimple1200series.SSLStatusEnabled),
				EndPoint:  to.Ptr("blob.core.windows.net"),
				Location:  to.Ptr("West US"),
				Login:     to.Ptr("SacForSDKTest"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple1200series%2Farmstorsimple1200series%2Fv1.0.0/sdk/resourcemanager/storsimple1200series/armstorsimple1200series/README.md) on how to add the SDK to your project and authenticate.
