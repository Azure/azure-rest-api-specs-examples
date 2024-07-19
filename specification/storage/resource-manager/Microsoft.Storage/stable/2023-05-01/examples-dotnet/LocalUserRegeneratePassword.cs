using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserRegeneratePassword.json
// this example is just showing the usage of "LocalUsers_RegeneratePassword" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountLocalUserResource created on azure
// for more information of creating StorageAccountLocalUserResource, please refer to the document of StorageAccountLocalUserResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string accountName = "sto2527";
string username = "user1";
ResourceIdentifier storageAccountLocalUserResourceId = StorageAccountLocalUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, username);
StorageAccountLocalUserResource storageAccountLocalUser = client.GetStorageAccountLocalUserResource(storageAccountLocalUserResourceId);

// invoke the operation
LocalUserRegeneratePasswordResult result = await storageAccountLocalUser.RegeneratePasswordAsync();

Console.WriteLine($"Succeeded: {result}");
