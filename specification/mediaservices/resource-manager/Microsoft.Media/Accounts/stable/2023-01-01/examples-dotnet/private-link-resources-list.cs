using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/private-link-resources-list.json
// this example is just showing the usage of "PrivateLinkResources_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountResource created on azure
// for more information of creating MediaServicesAccountResource, please refer to the document of MediaServicesAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosorg";
string accountName = "contososports";
ResourceIdentifier mediaServicesAccountResourceId = MediaServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MediaServicesAccountResource mediaServicesAccount = client.GetMediaServicesAccountResource(mediaServicesAccountResourceId);

// get the collection of this MediaServicesPrivateLinkResource
MediaServicesPrivateLinkResourceCollection collection = mediaServicesAccount.GetMediaServicesPrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (MediaServicesPrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MediaServicesPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
