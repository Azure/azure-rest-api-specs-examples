package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementServiceGetNetworkStatusByLocation.json
func ExampleNetworkStatusClient_ListByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkStatusClient().ListByLocation(ctx, "rg1", "apimService1", "North Central US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkStatusContract = armapimanagement.NetworkStatusContract{
	// 	ConnectivityStatus: []*armapimanagement.ConnectivityStatusContract{
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.blob.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.936Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:14.703Z"); return t}()),
	// 			ResourceType: to.Ptr("BlobStorage"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.file.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.926Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:41.532Z"); return t}()),
	// 			ResourceType: to.Ptr("FileStorage"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.queue.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.841Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:30.645Z"); return t}()),
	// 			ResourceType: to.Ptr("Queue"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.table.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.936Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:23.878Z"); return t}()),
	// 			ResourceType: to.Ptr("TableStorage"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T08:07:37.548Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:57:34.866Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://gcs.ppe.warm.ingestion.monitoring.azure.com"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.106Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:26.187Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://global.metrics.nsatc.net/"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.051Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:35.962Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://login.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.106Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:30.804Z"); return t}()),
	// 			ResourceType: to.Ptr("AzureActiveDirectory"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://prod2.metrics.nsatc.net:1886/RecoveryService"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.279Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:45.209Z"); return t}()),
	// 			ResourceType: to.Ptr("Metrics"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("LocalGatewayRedis"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.936Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:15.134Z"); return t}()),
	// 			ResourceType: to.Ptr("InternalCache"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.841Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:57.899Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("Scm"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:57.325Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T23:03:57.618Z"); return t}()),
	// 			ResourceType: to.Ptr("SourceControl"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("smtpi-xxx.msn.com:25028"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.351Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:58:22.243Z"); return t}()),
	// 			ResourceType: to.Ptr("Email"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("zwcvuxxxx.database.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.041Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:44.358Z"); return t}()),
	// 			ResourceType: to.Ptr("SQLDatabase"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 	}},
	// 	DNSServers: []*string{
	// 		to.Ptr("10.82.98.10")},
	// 	}
}
