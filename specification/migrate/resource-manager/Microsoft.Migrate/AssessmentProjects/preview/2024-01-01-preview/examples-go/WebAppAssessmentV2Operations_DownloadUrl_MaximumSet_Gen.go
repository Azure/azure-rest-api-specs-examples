package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/WebAppAssessmentV2Operations_DownloadUrl_MaximumSet_Gen.json
func ExampleWebAppAssessmentV2OperationsClient_BeginDownloadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebAppAssessmentV2OperationsClient().BeginDownloadURL(ctx, "rgopenapi", "sumukk-ccy-bcs4557project", "anraghun-selfhost-v2", "anraghun-selfhost-v2", map[string]any{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DownloadURL = armmigrationassessment.DownloadURL{
	// 	AssessmentReportURL: to.Ptr("https://assessmentsrvcanccysa2.blob.core.windows.net/cc1325ff-d42f-4c45-bfb6-12069e45becb/anraghun-selfhost-v2anraghun-v2-testWebAppAssessment.xlsx?sv=2018-03-28&sr=b&sig=NYBG6gJmofvIQsk1K1tewWfX51BpFZWsvIy7gBXNYUE%3D&st=2023-11-03T05%3A51%3A39Z&se=2023-11-03T06%3A26%3A39Z&sp=r"),
	// 	ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:09.465Z"); return t}()),
	// }
}
