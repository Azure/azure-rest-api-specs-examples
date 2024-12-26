using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBSqlDatabaseThroughputGet.json
// this example is just showing the usage of "SqlResources_GetSqlDatabaseThroughput" operation, for the dependent resources, they will have to be created separately.

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
CosmosDBSqlDatabaseThroughputSettingResource result = await cosmosDBSqlDatabaseThroughputSetting.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ThroughputSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
