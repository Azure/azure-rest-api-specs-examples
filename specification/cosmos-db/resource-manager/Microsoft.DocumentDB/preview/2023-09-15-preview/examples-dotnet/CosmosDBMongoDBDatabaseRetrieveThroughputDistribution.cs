using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBMongoDBDatabaseRetrieveThroughputDistribution.json
// this example is just showing the usage of "MongoDBResources_MongoDBDatabaseRetrieveThroughputDistribution" operation, for the dependent resources, they will have to be created separately.

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
ArmOperation<PhysicalPartitionThroughputInfoResult> lro = await mongoDBDatabaseThroughputSetting.MongoDBDatabaseRetrieveThroughputDistributionAsync(WaitUntil.Completed, retrieveThroughputParameters);
PhysicalPartitionThroughputInfoResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
