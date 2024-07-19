using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBMongoDBDatabaseMigrateToAutoscale.json
// this example is just showing the usage of "MongoDBResources_MigrateMongoDBDatabaseToAutoscale" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBDatabaseThroughputSettingResource created on azure
// for more information of creating MongoDBDatabaseThroughputSettingResource, please refer to the document of MongoDBDatabaseThroughputSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
ResourceIdentifier mongoDBDatabaseThroughputSettingResourceId = MongoDBDatabaseThroughputSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
MongoDBDatabaseThroughputSettingResource mongoDBDatabaseThroughputSetting = client.GetMongoDBDatabaseThroughputSettingResource(mongoDBDatabaseThroughputSettingResourceId);

// invoke the operation
ArmOperation<MongoDBDatabaseThroughputSettingResource> lro = await mongoDBDatabaseThroughputSetting.MigrateMongoDBDatabaseToAutoscaleAsync(WaitUntil.Completed);
MongoDBDatabaseThroughputSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ThroughputSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
