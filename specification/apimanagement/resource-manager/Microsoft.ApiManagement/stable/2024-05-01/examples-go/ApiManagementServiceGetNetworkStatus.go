package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementServiceGetNetworkStatus.json
func ExampleNetworkStatusClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkStatusClient().ListByService(ctx, "rg1", "apimService1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkStatusContractByLocationArray = []*armapimanagement.NetworkStatusContractByLocation{
	// 	{
	// 		Location: to.Ptr("Central US EUAP"),
	// 		NetworkStatus: &armapimanagement.NetworkStatusContract{
	// 			ConnectivityStatus: []*armapimanagement.ConnectivityStatusContract{
	// 				{
	// 					Name: to.Ptr("apimst8rlxXXXXX7.queue.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.178Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:41:44.639Z"); return t}()),
	// 					ResourceType: to.Ptr("SmtpQueue"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimstm5lkXXXXX4kky0.blob.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.178Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:39:33.942Z"); return t}()),
	// 					ResourceType: to.Ptr("BlobStorage"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimstm5lkXXXXX4kky0.file.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:43.895Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:41:38.999Z"); return t}()),
	// 					ResourceType: to.Ptr("FileStorage"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimstm5lkXXXXX4kky0.queue.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.208Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:41:38.436Z"); return t}()),
	// 					ResourceType: to.Ptr("Queue"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimstm5lkXXXXX4kky0.table.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:43.973Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:41:17.476Z"); return t}()),
	// 					ResourceType: to.Ptr("TableStorage"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apirpsqlmgsXXXXXXic45.database.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.070Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:50.295Z"); return t}()),
	// 					ResourceType: to.Ptr("SQLDatabase"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:43.848Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:09.844Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://apikv-XXXXXXurugl3.vault.azure.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-02T06:54:45.712Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:17.344Z"); return t}()),
	// 					ResourceType: to.Ptr("AzureKeyVault"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://centraluseuap.login.microsoft.com"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.149Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:01.666Z"); return t}()),
	// 					ResourceType: to.Ptr("RegionalAzureActiveDirectory"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://gcs.prod.warm.ingestion.monitoring.azure.com"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-26T14:09:26.250Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:42:08.005Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://global.prod.microsoftmetrics.com/"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.129Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:12.969Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://login.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-04T11:10:16.104Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:43.404Z"); return t}()),
	// 					ResourceType: to.Ptr("AzureActiveDirectory"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://partner.prod.repmap.microsoft.com"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.208Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:40:32.606Z"); return t}()),
	// 					ResourceType: to.Ptr("CaptchaEndpoint"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://xxxxx.prod.microsoftmetrics.com:1886/RecoveryService"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.348Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:42:33.856Z"); return t}()),
	// 					ResourceType: to.Ptr("Metrics"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://samir-XXXXXXX.management.azure-api.net:3443/servicestatus?api-version=2018-01-01"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:42:10.395Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:42:10.395Z"); return t}()),
	// 					ResourceType: to.Ptr("ApiManagement Control Plane - inbound"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("LocalGatewayRedis"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:53:44.208Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:39:30.184Z"); return t}()),
	// 					ResourceType: to.Ptr("InternalCache"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("Scm"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-25T03:32:02.648Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:41:24.662Z"); return t}()),
	// 					ResourceType: to.Ptr("SourceControl"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 			}},
	// 			DNSServers: []*string{
	// 				to.Ptr("168.63.129.16")},
	// 			},
	// 	}}
}
