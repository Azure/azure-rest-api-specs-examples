package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Common/ExportJobsOperationResult.json
func ExampleExportJobsOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportJobsOperationResultsClient().Get(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationResultInfoBaseResource = armrecoveryservicesbackup.OperationResultInfoBaseResource{
	// 	Headers: map[string][]*string{
	// 	},
	// 	Operation: &armrecoveryservicesbackup.ExportJobsOperationResultInfo{
	// 		ObjectType: to.Ptr("ExportJobsOperationResultInfo"),
	// 		BlobSasKey: to.Ptr("?sv=2014-02-14&sr=b&sig=<sas_signature>&st=2017-11-29T07%3A53%3A34Z&se=2017-11-29T08%3A03%3A34Z&sp=r"),
	// 		BlobURL: to.Ptr("https://azureblob.blob.core.windows.net/reportcontainer/exportjobsreportc00000000-0000-0000-0000-000000000000"),
	// 	},
	// }
}
