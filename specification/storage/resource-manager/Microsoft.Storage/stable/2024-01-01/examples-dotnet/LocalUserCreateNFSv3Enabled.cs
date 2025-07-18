using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/LocalUserCreateNFSv3Enabled.json
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
StorageAccountLocalUserData data = new StorageAccountLocalUserData
{
    ExtendedGroups = { 1001, 1005, 2005 },
    IsNfsV3Enabled = true,
};
ArmOperation<StorageAccountLocalUserResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, username, data);
StorageAccountLocalUserResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountLocalUserData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
