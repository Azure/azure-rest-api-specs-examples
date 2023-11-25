package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Assessments_GetReportDownloadUrl.json
func ExampleAssessmentsClient_GetReportDownloadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().GetReportDownloadURL(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "Test1", "assessment_5_9_2019_16_22_14", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DownloadURL = armmigrate.DownloadURL{
	// 	AssessmentReportURL: to.Ptr("link-to-download-assessment-report"),
	// 	ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-05T13:17:23.543Z"); return t}()),
	// }
}
