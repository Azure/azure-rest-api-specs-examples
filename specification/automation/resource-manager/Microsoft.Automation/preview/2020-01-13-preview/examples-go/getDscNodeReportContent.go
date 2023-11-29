package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeReportContent.json
func ExampleNodeReportsClient_GetContent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNodeReportsClient().GetContent(ctx, "rg", "myAutomationAccount33", "nodeId", "reportId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = map[string]any{
	// 	"AdditionalData":[]any{
	// 		map[string]any{
	// 			"Key": "OSVersion",
	// 			"Value": "{\"VersionString\":\"Microsoft Windows NT 6.1.7601 Service Pack 1\",\"ServicePack\":\"Service Pack 1\",\"Platform\":\"Win32NT\"}",
	// 		},
	// 		map[string]any{
	// 			"Key": "PSVersion",
	// 			"Value": "{\"CLRVersion\":\"4.0.30319.42000\",\"PSVersion\":\"5.1.14409.1012\",\"BuildVersion\":\"10.0.14409.1012\"}",
	// 		},
	// 	},
	// 	"Errors":[]any{
	// 	},
	// 	"IpAddress": "10.13.49.8;127.0.0.1;fe80::2cc0:8062:a210:e1c6%11;::2000:0:0:0;::1;::2000:0:0:0",
	// 	"JobId": "eabe061f-2e1f-11e8-8d01-000d3a18dec4",
	// 	"LCMVersion": "2.0",
	// 	"NodeName": "ANAGG-2008R2",
	// 	"OperationType": "Consistency",
	// 	"ReportFormatVersion": "2.0",
	// 	"StartTime": "2018-03-22T22:25:26.2140000+00:00",
	// 	"StatusData":[]any{
	// 		"{\"IPV4Addresses\":[\"10.13.49.8\",\"127.0.0.1\"],\"MACAddresses\":[\"00-0D-3A-18-DE-C4\",\"00-00-00-00-00-00-00-E0\"],\"Type\":\"Consistency\",\"HostName\":\"ANAGG-2008R2\",\"Locale\":\"en-US\",\"StartDate\":\"2018-03-22T22:25:26.2140000+00:00\",\"JobID\":\"{EABE061F-2E1F-11E8-8D01-000D3A18DEC4}\",\"LCMVersion\":\"2.0\",\"IPV6Addresses\":[\"fe80::2cc0:8062:a210:e1c6%11\",\"::2000:0:0:0\",\"::1\",\"::2000:0:0:0\"]}",
	// 	},
	// }
}
