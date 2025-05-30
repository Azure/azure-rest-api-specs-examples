package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationListResult = armredis.OperationListResult{
		// 	Value: []*armredis.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/checknameavailability/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Checks if a name is available for use with a new Redis Cache"),
		// 				Operation: to.Ptr("Check Cache Name Availability"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/register/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Registers the 'Microsoft.Cache' resource provider with a subscription"),
		// 				Operation: to.Ptr("Register Resource Provider Microsoft.Cache"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/unregister/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Unregisters the 'Microsoft.Cache' resource provider with a subscription"),
		// 				Operation: to.Ptr("Unregister Resource Provider Microsoft.Cache"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/operations/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Lists the operations that 'Microsoft.Cache' provider supports."),
		// 				Operation: to.Ptr("List Provider Operations"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/locations/operationResults/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Gets the result of a long running operation for which the 'Location' header was previously returned to the client"),
		// 				Operation: to.Ptr("Read operation results"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/locations/operationsStatus/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the status of a long running operation for which the 'AzureAsync' header was previously returned to the client"),
		// 				Operation: to.Ptr("Read the status of a long running operation"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/locations/asyncOperations/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read an Async Operation's Status"),
		// 				Operation: to.Ptr("Read asynchronous operation status"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/locations/checknameavailability/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Checks if a name is available for use with a new Redis Enterprise cache"),
		// 				Operation: to.Ptr("Check Cache Name Availability in location"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Modify the Redis Cache's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Cache (read-write)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the Redis Cache's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Cache (read-only)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete the entire Redis Cache"),
		// 				Operation: to.Ptr("Delete Redis Cache"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/listKeys/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the value of Redis Cache access keys in the management portal"),
		// 				Operation: to.Ptr("View Redis Cache Access Keys"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/regenerateKey/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Change the value of Redis Cache access keys in the management portal"),
		// 				Operation: to.Ptr("Regenerate Redis Cache Access Keys"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/import/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Import data of a specified format from multiple blobs into Redis"),
		// 				Operation: to.Ptr("Import data into Redis from storage"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/export/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Export Redis data to prefixed storage blobs in specified format"),
		// 				Operation: to.Ptr("Export Redis data to storage"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/forceReboot/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Force reboot a cache instance, potentially with data loss."),
		// 				Operation: to.Ptr("Force reboot a cache instance, potentially with data loss."),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/stop/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Stop an Azure Cache for Redis, potentially with data loss."),
		// 				Operation: to.Ptr("Stop an Azure Cache for Redis, potentially with data loss."),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/start/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Start an Azure Cache for Redis"),
		// 				Operation: to.Ptr("Start an Azure Cache for Redis"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete the entire Redis Enterprise cache"),
		// 				Operation: to.Ptr("Delete Redis Enterprise cache"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the Redis Enterprise cache's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Enterprise cache (read)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Modify the Redis Enterprise cache's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Enterprise cache (write)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Deletes a Redis Enterprise database and its contents"),
		// 				Operation: to.Ptr("Delete Redis Enterprise database"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the Redis Enterprise cache database's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Enterprise cache database (read)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Modify the Redis Enterprise cache database's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Enterprise cache database (write)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/export/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Export data to storage blobs from a Redis Enterprise database "),
		// 				Operation: to.Ptr("Export Redis Enterprise database"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/forceUnlink/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Forcibly unlink a georeplica Redis Enterprise database from its peers"),
		// 				Operation: to.Ptr("Force unlink Redis Enterprise database georeplica"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/import/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Import data from storage blobs to a Redis Enterprise database"),
		// 				Operation: to.Ptr("Import Redis Enterprise database"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/listKeys/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the value of Redis Enterprise database access keys in the management portal"),
		// 				Operation: to.Ptr("View Redis Enterprise database access keys"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/regenerateKey/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Change the value of Redis Enterprise database access keys in the management portal"),
		// 				Operation: to.Ptr("Regenerate Redis Enterprise database access keys"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/databases/operationResults/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the result of Redis Enterprise database operations in the management portal"),
		// 				Operation: to.Ptr("View Redis Enterprise database operation results"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise database operation results"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/operationResults/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the result of Redis Enterprise operations in the management portal"),
		// 				Operation: to.Ptr("View Redis Enterprise operation results"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise operation results"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/metricDefinitions/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for a Redis Cache"),
		// 				Operation: to.Ptr("Read Redis Cache Metric Definitions"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("The available metrics for a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for a Redis Enterprise Cache"),
		// 				Operation: to.Ptr("Read Redis Enterprise Metric Definitions"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("The available metrics for a Redis Enterprise Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/patchSchedules/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Gets the patching schedule of a Redis Cache"),
		// 				Operation: to.Ptr("Get Redis Cache Patch Schedule"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Patching schedule of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/patchSchedules/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Modify the patching schedule of a Redis Cache"),
		// 				Operation: to.Ptr("Change Redis Patching Schedule"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Patching schedule of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/patchSchedules/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete the patch schedule of a Redis Cache"),
		// 				Operation: to.Ptr("Delete Redis Cache Patch Schedule"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Patching schedule of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/firewallRules/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Get the IP firewall rules of a Redis Cache"),
		// 				Operation: to.Ptr("Get Redis Cache Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("IP firewall rule of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/firewallRules/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Edit the IP firewall rules of a Redis Cache"),
		// 				Operation: to.Ptr("Update Redis Cache Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("IP firewall rule of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/firewallRules/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete IP firewall rules of a Redis Cache"),
		// 				Operation: to.Ptr("Delete Redis Cache Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("IP firewall rule of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/linkedServers/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Get Linked Servers associated with a redis cache."),
		// 				Operation: to.Ptr("Get Redis Cache Linked Servers"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Linked Servers of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/linkedServers/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Add Linked Server to a Redis Cache"),
		// 				Operation: to.Ptr("Add Redis Cache Linked Server"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Linked Servers of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/linkedServers/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete Linked Server from a Redis Cache"),
		// 				Operation: to.Ptr("Delete Redis Cache Linked Server"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Linked Servers of a Redis Cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/eventGridFilters/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Get Redis Cache Event Grid Filter"),
		// 				Operation: to.Ptr("Get Redis Cache Event Grid Filter"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache Event Grid Filter"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/eventGridFilters/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Update Redis Cache Event Grid Filters"),
		// 				Operation: to.Ptr("Update Redis Cache Event Grid Filters"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache Event Grid Filter"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/eventGridFilters/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete Redis Cache Event Grid Filters"),
		// 				Operation: to.Ptr("Delete Redis Cache Event Grid Filters"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Cache Event Grid Filter"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Validate the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Validate private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnectionProxies/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Get the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Get private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnectionProxies/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Create the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Create private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnectionProxies/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Delete private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Validate the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Validate private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnectionProxies/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Get the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Get private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnectionProxies/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Create the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Create private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnectionProxies/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete the private endpoint connection proxy"),
		// 				Operation: to.Ptr("Delete private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnectionProxies/operationResults/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("View the result of private endpoint connection operations in the management portal"),
		// 				Operation: to.Ptr("Redis Enterprise cache private endpoint operation results (read)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise private endpoint connection proxies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnections/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read a private endpoint connection"),
		// 				Operation: to.Ptr("Read private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connections"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnections/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Write a private endpoint connection"),
		// 				Operation: to.Ptr("Write private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connections"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateEndpointConnections/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete a private endpoint connection"),
		// 				Operation: to.Ptr("Delete private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private endpoint connections"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read a private endpoint connection"),
		// 				Operation: to.Ptr("Read private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache private endpoint connections"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Write a private endpoint connection"),
		// 				Operation: to.Ptr("Write private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache private endpoint connections"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete a private endpoint connection"),
		// 				Operation: to.Ptr("Delete private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache private endpoint connections"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/privateLinkResources/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read 'groupId' of redis subresource that a private link can be connected to"),
		// 				Operation: to.Ptr("Read Private Linkable Resources"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private linkable resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/privateLinkResources/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read 'groupId' of redis subresource that a private link can be connected to"),
		// 				Operation: to.Ptr("Read Private Linkable Resources"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache private link resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/roles/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read roles on a Redis Cache"),
		// 				Operation: to.Ptr("Read Redis Roles"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis Roles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/roles/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Create or update role on a Redis Cache"),
		// 				Operation: to.Ptr("Update Redis Roles"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis Roles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/roles/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete role on a Redis Cache"),
		// 				Operation: to.Ptr("Delete Redis Roles"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis Roles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/roleAssignments/read"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Read role assignments on a Redis Cache"),
		// 				Operation: to.Ptr("Read Redis Role Assignments"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis Role Description"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/roleAssignments/write"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Create or update role assignments on a Redis Cache"),
		// 				Operation: to.Ptr("Update Redis Role Assignments"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis Role Description"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/roleAssignments/delete"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Delete All Redis Role Assignments"),
		// 				Operation: to.Ptr("Delete Redis Role Assignment"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis Role Description"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redis/PrivateEndpointConnectionsApproval/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Approve Private Endpoint Connections"),
		// 				Operation: to.Ptr("Approve Private Endpoint Connections"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Azure Cache for Redis private linkable resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/PrivateEndpointConnectionsApproval/action"),
		// 			Display: &armredis.OperationDisplay{
		// 				Description: to.Ptr("Approve Private Endpoint Connections"),
		// 				Operation: to.Ptr("Approve Private Endpoint Connections"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache private link resources"),
		// 			},
		// 	}},
		// }
	}
}
