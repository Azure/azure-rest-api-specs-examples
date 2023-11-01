using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/mongo-cluster/CosmosDBMongoClusterUpdate.json
// this example is just showing the usage of "MongoClusters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoClusterResource created on azure
// for more information of creating MongoClusterResource, please refer to the document of MongoClusterResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestResourceGroup";
string mongoClusterName = "myMongoCluster";
ResourceIdentifier mongoClusterResourceId = MongoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mongoClusterName);
MongoClusterResource mongoCluster = client.GetMongoClusterResource(mongoClusterResourceId);

// invoke the operation
MongoClusterPatch patch = new MongoClusterPatch()
{
    AdministratorLogin = "mongoAdmin",
    AdministratorLoginPassword = "password",
    ServerVersion = "5.0",
    NodeGroupSpecs =
    {
    new NodeGroupSpec()
    {
    Kind = NodeKind.Shard,
    NodeCount = 4,
    Sku = "M50",
    DiskSizeInGB = 256,
    EnableHa = true,
    }
    },
};
ArmOperation<MongoClusterResource> lro = await mongoCluster.UpdateAsync(WaitUntil.Completed, patch);
MongoClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
