using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MongoCluster.Models;
using Azure.ResourceManager.MongoCluster;

// Generated from example definition: 2025-09-01/MongoClusters_UserCreateOrUpdate.json
// this example is just showing the usage of "User_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoClusterUserResource created on azure
// for more information of creating MongoClusterUserResource, please refer to the document of MongoClusterUserResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string mongoClusterName = "myMongoCluster";
string userName = "uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu";
ResourceIdentifier mongoClusterUserResourceId = MongoClusterUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mongoClusterName, userName);
MongoClusterUserResource mongoClusterUser = client.GetMongoClusterUserResource(mongoClusterUserResourceId);

// invoke the operation
MongoClusterUserData data = new MongoClusterUserData
{
    Properties = new MongoClusterUserProperties
    {
        IdentityProvider = new MongoClusterEntraIdentityProvider(new MongoClusterEntraIdentityProviderProperties(MongoClusterEntraPrincipalType.User)),
        Roles = { new MongoClusterDatabaseRole("admin", MongoClusterUserRole.Root) },
    },
};
ArmOperation<MongoClusterUserResource> lro = await mongoClusterUser.UpdateAsync(WaitUntil.Completed, data);
MongoClusterUserResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoClusterUserData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
