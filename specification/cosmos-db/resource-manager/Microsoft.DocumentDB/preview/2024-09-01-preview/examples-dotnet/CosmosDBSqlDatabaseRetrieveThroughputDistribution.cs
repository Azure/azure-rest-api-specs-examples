using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlDatabaseRetrieveThroughputDistribution.json
// this example is just showing the usage of "SqlResources_SqlDatabaseRetrieveThroughputDistribution" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlDatabaseThroughputSettingResource created on azure
// for more information of creating CosmosDBSqlDatabaseThroughputSettingResource, please refer to the document of CosmosDBSqlDatabaseThroughputSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
ResourceIdentifier cosmosDBSqlDatabaseThroughputSettingResourceId = CosmosDBSqlDatabaseThroughputSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
CosmosDBSqlDatabaseThroughputSettingResource cosmosDBSqlDatabaseThroughputSetting = client.GetCosmosDBSqlDatabaseThroughputSettingResource(cosmosDBSqlDatabaseThroughputSettingResourceId);

// invoke the operation
RetrieveThroughputParameters retrieveThroughputParameters = new RetrieveThroughputParameters(new AzureLocation("placeholder"), new RetrieveThroughputPropertiesResource(new WritableSubResource[]
{
new WritableSubResource()
{
Id = new ResourceIdentifier("0"),
},new WritableSubResource()
{
Id = new ResourceIdentifier("1"),
}
}));
ArmOperation<PhysicalPartitionThroughputInfoResult> lro = await cosmosDBSqlDatabaseThroughputSetting.SqlDatabaseRetrieveThroughputDistributionAsync(WaitUntil.Completed, retrieveThroughputParameters);
PhysicalPartitionThroughputInfoResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
