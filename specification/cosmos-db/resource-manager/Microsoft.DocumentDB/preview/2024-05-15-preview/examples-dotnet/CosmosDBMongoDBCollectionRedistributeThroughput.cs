using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBMongoDBCollectionRedistributeThroughput.json
// this example is just showing the usage of "MongoDBResources_MongoDBContainerRedistributeThroughput" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBCollectionThroughputSettingResource created on azure
// for more information of creating MongoDBCollectionThroughputSettingResource, please refer to the document of MongoDBCollectionThroughputSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string collectionName = "collectionName";
ResourceIdentifier mongoDBCollectionThroughputSettingResourceId = MongoDBCollectionThroughputSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, collectionName);
MongoDBCollectionThroughputSettingResource mongoDBCollectionThroughputSetting = client.GetMongoDBCollectionThroughputSettingResource(mongoDBCollectionThroughputSettingResourceId);

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
ArmOperation<PhysicalPartitionThroughputInfoResult> lro = await mongoDBCollectionThroughputSetting.MongoDBContainerRedistributeThroughputAsync(WaitUntil.Completed, redistributeThroughputParameters);
PhysicalPartitionThroughputInfoResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
