using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MongoCluster.Models;
using Azure.ResourceManager.MongoCluster;

// Generated from example definition: 2025-09-01/MongoClusters_PatchCMK.json
// this example is just showing the usage of "MongoCluster_Update" operation, for the dependent resources, they will have to be created separately.

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
MongoClusterPatch patch = new MongoClusterPatch
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity2")] = new UserAssignedIdentity()
        },
    },
    Properties = new MongoClusterUpdateProperties
    {
        CustomerManagedKeyEncryption = new MongoClusterCmkEncryptionProperties
        {
            KeyEncryptionKeyIdentity = new MongoClusterKeyEncryptionKeyIdentity
            {
                IdentityType = MongoClusterKeyEncryptionKeyIdentityType.UserAssignedIdentity,
                UserAssignedIdentityResourceId = "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity2",
            },
            KeyEncryptionKeyUri = "https://myVault.vault.azure.net/keys/myKey2",
        },
    },
};
ArmOperation<MongoClusterResource> lro = await mongoCluster.UpdateAsync(WaitUntil.Completed, patch);
MongoClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
