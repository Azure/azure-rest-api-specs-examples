package armspringappdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/ErrorSummaries_Get_MaximumSet_Gen.json
func ExampleErrorSummariesClient_Get_errorSummariesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewErrorSummariesClient().Get(ctx, "rgspringbootdiscovery", "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps", "K2lv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ErrorSummary = armspringappdiscovery.ErrorSummary{
	// 	Name: to.Ptr("tfmbwokvgn"),
	// 	Type: to.Ptr("qfemnwktjpynmezppab"),
	// 	ID: to.Ptr("waudmuluqttorwxywyibbezvko"),
	// 	SystemData: &armspringappdiscovery.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T16:44:32.561Z"); return t}()),
	// 		CreatedBy: to.Ptr("ztjtyfhicmxcpqszeovgojwzceagbc"),
	// 		CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T16:44:32.562Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("isjllzvqodp"),
	// 		LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
	// 	},
	// 	Properties: &armspringappdiscovery.ErrorSummariesProperties{
	// 		DiscoveryScopeErrorSummaries: []*armspringappdiscovery.ErrorSummaryModel{
	// 			{
	// 				AffectedObjectsCount: to.Ptr[int64](2),
	// 				AffectedResourceType: to.Ptr("kprvjasvybficrqctgbjpaek"),
	// 		}},
	// 		Errors: []*armspringappdiscovery.Error{
	// 			{
	// 				Code: to.Ptr("czbrpdxmv"),
	// 				ID: to.Ptr[int64](13),
	// 				Message: to.Ptr("knjufnfkdpukqiuqzfviwnss"),
	// 				PossibleCauses: to.Ptr("knjufnfkdpukqiuqzfviwnss"),
	// 				RecommendedAction: to.Ptr("qpycieevlbrcomlwooiw"),
	// 				RunAsAccountID: to.Ptr("knjufnfkdpukqiuqzfviwnss"),
	// 				Severity: to.Ptr("wcusqqmqwo"),
	// 				SummaryMessage: to.Ptr("knjufnfkdpukqiuqzfviwnss"),
	// 				UpdatedTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 		}},
	// 		ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"key2085": to.Ptr("olljrx"),
	// 	},
	// }
}
