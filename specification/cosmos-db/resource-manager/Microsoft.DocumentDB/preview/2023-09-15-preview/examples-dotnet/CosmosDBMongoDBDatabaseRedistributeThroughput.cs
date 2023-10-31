using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBMongoDBDatabaseRedistributeThroughput.json
// this example is just showing the usage of "MongoDBResources_MongoDBDatabaseRedistributeThroughput" operation, for the dependent resources, they will have to be created separately.

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
RedistributeThroughputParameters redistributeThroughputParameters = new RedistributeThroughputParameters(new AzureLocation("placeholder"), new RedistributeThroughputPropertiesResource(ThroughputPolicyType.Custom, new PhysicalPartitionThroughputInfoResource[]
{
new PhysicalPartitionThroughputInfoResource("0")
{
Throughput = 5000,
},new PhysicalPartitionThroughputInfoResource("1")
{
Throughput = 5000,
}
}, new PhysicalPartitionThroughputInfoResource[]
{
new PhysicalPartitionThroughputInfoResource("2")
{
Throughput = 5000,
},new PhysicalPartitionThroughputInfoResource("3")
}));
ArmOperation<PhysicalPartitionThroughputInfoResult> lro = await mongoDBDatabaseThroughputSetting.MongoDBDatabaseRedistributeThroughputAsync(WaitUntil.Completed, redistributeThroughputParameters);
PhysicalPartitionThroughputInfoResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
