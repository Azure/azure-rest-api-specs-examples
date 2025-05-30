using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBPKeyRangeIdRegionGetMetrics.json
// this example is just showing the usage of "PartitionKeyRangeIdRegion_ListMetrics" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBAccountResource created on azure
// for more information of creating CosmosDBAccountResource, please refer to the document of CosmosDBAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
ResourceIdentifier cosmosDBAccountResourceId = CosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CosmosDBAccountResource cosmosDBAccount = client.GetCosmosDBAccountResource(cosmosDBAccountResourceId);

// invoke the operation and iterate over the result
string region = "West US";
string databaseRid = "databaseRid";
string collectionRid = "collectionRid";
string partitionKeyRangeId = "0";
string filter = "$filter=(name.value eq 'Max RUs Per Second') and timeGrain eq duration'PT1M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T23:58:55.2780000Z";
await foreach (PartitionMetric item in cosmosDBAccount.GetMetricsPartitionKeyRangeIdRegionsAsync(region, databaseRid, collectionRid, partitionKeyRangeId, filter))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
