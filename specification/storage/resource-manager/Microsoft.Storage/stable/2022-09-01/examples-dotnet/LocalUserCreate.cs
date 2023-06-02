using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/LocalUserCreate.json
// this example is just showing the usage of "LocalUsers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string accountName = "sto2527";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// get the collection of this StorageAccountLocalUserResource
StorageAccountLocalUserCollection collection = storageAccount.GetStorageAccountLocalUsers();

// invoke the operation
string username = "user1";
StorageAccountLocalUserData data = new StorageAccountLocalUserData()
{
    PermissionScopes =
    {
    new StoragePermissionScope("rwd","file","share1"),new StoragePermissionScope("rw","file","share2")
    },
    HomeDirectory = "homedirectory",
    SshAuthorizedKeys =
    {
    new StorageSshPublicKey()
    {
    Description = "key name",
    Key = "ssh-rsa keykeykeykeykey=",
    }
    },
    HasSshPassword = true,
};
ArmOperation<StorageAccountLocalUserResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, username, data);
StorageAccountLocalUserResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountLocalUserData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
