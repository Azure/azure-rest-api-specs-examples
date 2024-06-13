package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/incidents/GetAllIncidentBookmarks.json
func ExampleIncidentsClient_ListBookmarks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().ListBookmarks(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IncidentBookmarkList = armsecurityinsights.IncidentBookmarkList{
	// 	Value: []*armsecurityinsights.HuntingBookmark{
	// 		{
	// 			Name: to.Ptr("afbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/afbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindBookmark),
	// 			Properties: &armsecurityinsights.HuntingBookmarkProperties{
	// 				AdditionalData: map[string]any{
	// 					"ETag": "\"3b00acab-0000-0d00-0000-5f15e4ed0000\"",
	// 					"EntityId": "afbd324f-6c48-459c-8710-8d1e1cd03812",
	// 				},
	// 				FriendlyName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				CreatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("b03ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 				DisplayName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				Labels: []*string{
	// 				},
	// 				Query: to.Ptr("SecurityEvent\r\n| take 1\n"),
	// 				QueryResult: to.Ptr("{\"TimeGenerated\":\"2020-05-24T01:24:25.67Z\",\"Account\":\"\\\\ADMINISTRATOR\",\"AccountType\":\"User\",\"Computer\":\"SecurityEvents\",\"EventSourceName\":\"Microsoft-Windows-Security-Auditing\",\"Channel\":\"Security\",\"Task\":12544,\"Level\":\"16\",\"EventID\":4625,\"Activity\":\"4625 - An account failed to log on.\",\"AuthenticationPackageName\":\"NTLM\",\"FailureReason\":\"%%2313\",\"IpAddress\":\"176.113.115.73\",\"IpPort\":\"0\",\"LmPackageName\":\"-\",\"LogonProcessName\":\"NtLmSsp \",\"LogonType\":3,\"LogonTypeName\":\"3 - Network\",\"Process\":\"-\",\"ProcessId\":\"0x0\",\"__entityMapping\":{\"\\\\ADMINISTRATOR\":\"Account\",\"SecurityEvents\":\"Host\"}}"),
	// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				UpdatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("b03ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("bbbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/bbbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindBookmark),
	// 			Properties: &armsecurityinsights.HuntingBookmarkProperties{
	// 				AdditionalData: map[string]any{
	// 					"ETag": "\"3b00acab-0000-0d00-0000-5f15e4ed0000\"",
	// 					"EntityId": "afbd324f-6c48-459c-8710-8d1e1cd03812",
	// 				},
	// 				FriendlyName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				CreatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("303ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 				DisplayName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				Labels: []*string{
	// 				},
	// 				Query: to.Ptr("SecurityEvent\r\n| take 1\n"),
	// 				QueryResult: to.Ptr("{\"TimeGenerated\":\"2020-05-24T01:24:25.67Z\",\"Account\":\"\\\\ADMINISTRATOR\",\"AccountType\":\"User\",\"Computer\":\"SecurityEvents\",\"EventSourceName\":\"Microsoft-Windows-Security-Auditing\",\"Channel\":\"Security\",\"Task\":12544,\"Level\":\"16\",\"EventID\":4625,\"Activity\":\"4625 - An account failed to log on.\",\"AuthenticationPackageName\":\"NTLM\",\"FailureReason\":\"%%2313\",\"IpAddress\":\"176.113.115.73\",\"IpPort\":\"0\",\"LmPackageName\":\"-\",\"LogonProcessName\":\"NtLmSsp \",\"LogonType\":3,\"LogonTypeName\":\"3 - Network\",\"Process\":\"-\",\"ProcessId\":\"0x0\",\"__entityMapping\":{\"\\\\ADMINISTRATOR\":\"Account\",\"SecurityEvents\":\"Host\"}}"),
	// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				UpdatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("b03ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 			},
	// 	}},
	// }
}
