using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accounts-sync-storage-keys.json
// this example is just showing the usage of "Mediaservices_SyncStorageKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountResource created on azure
// for more information of creating MediaServicesAccountResource, please refer to the document of MediaServicesAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contososports";
ResourceIdentifier mediaServicesAccountResourceId = MediaServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MediaServicesAccountResource mediaServicesAccount = client.GetMediaServicesAccountResource(mediaServicesAccountResourceId);

// invoke the operation
SyncStorageKeysContent content = new SyncStorageKeysContent()
{
    Id = "contososportsstore",
};
await mediaServicesAccount.SyncStorageKeysAsync(content);

Console.WriteLine($"Succeeded");
