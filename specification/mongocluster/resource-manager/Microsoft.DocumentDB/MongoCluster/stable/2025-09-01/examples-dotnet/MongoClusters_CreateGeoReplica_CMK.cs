using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MongoCluster.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MongoCluster;

// Generated from example definition: 2025-09-01/MongoClusters_CreateGeoReplica_CMK.json
// this example is just showing the usage of "MongoCluster_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
string mongoClusterName = "myReplicaMongoCluster";
MongoClusterData data = new MongoClusterData(new AzureLocation("centralus"))
{
    Properties = new MongoClusterProperties
    {
        CreateMode = MongoClusterCreateMode.GeoReplica,
        ReplicaParameters = new MongoClusterReplicaContent(new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/mySourceMongoCluster"), new AzureLocation("eastus")),
        CustomerManagedKeyEncryption = new MongoClusterCmkEncryptionProperties
        {
            KeyEncryptionKeyIdentity = new MongoClusterKeyEncryptionKeyIdentity
            {
                IdentityType = MongoClusterKeyEncryptionKeyIdentityType.UserAssignedIdentity,
                UserAssignedIdentityResourceId = "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
            },
            KeyEncryptionKeyUri = "https://myVault.vault.azure.net/keys/myKey",
        },
    },
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity")] = new UserAssignedIdentity()
        },
    },
};
ArmOperation<MongoClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, mongoClusterName, data);
MongoClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
