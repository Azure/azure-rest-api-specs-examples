using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/mongo-cluster/CosmosDBMongoClusterCreate.json
// this example is just showing the usage of "MongoClusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MongoClusterResource
MongoClusterCollection collection = resourceGroupResource.GetMongoClusters();

// invoke the operation
string mongoClusterName = "myMongoCluster";
MongoClusterData data = new MongoClusterData(new AzureLocation("westus2"))
{
    AdministratorLogin = "mongoAdmin",
    AdministratorLoginPassword = "password",
    ServerVersion = "5.0",
    NodeGroupSpecs =
    {
    new NodeGroupSpec()
    {
    Kind = NodeKind.Shard,
    NodeCount = 3,
    Sku = "M30",
    DiskSizeInGB = 128,
    EnableHa = true,
    }
    },
};
ArmOperation<MongoClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, mongoClusterName, data);
MongoClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
