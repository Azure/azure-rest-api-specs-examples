using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBMongoDBCollectionRetrieveThroughputDistribution.json
// this example is just showing the usage of "MongoDBResources_MongoDBContainerRetrieveThroughputDistribution" operation, for the dependent resources, they will have to be created separately.

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
ArmOperation<PhysicalPartitionThroughputInfoResult> lro = await mongoDBCollectionThroughputSetting.MongoDBContainerRetrieveThroughputDistributionAsync(WaitUntil.Completed, retrieveThroughputParameters);
PhysicalPartitionThroughputInfoResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
