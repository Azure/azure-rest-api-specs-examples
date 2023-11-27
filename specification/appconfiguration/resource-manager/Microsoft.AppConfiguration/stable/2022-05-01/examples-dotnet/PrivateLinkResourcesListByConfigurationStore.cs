using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppConfiguration;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/PrivateLinkResourcesListByConfigurationStore.json
// this example is just showing the usage of "PrivateLinkResources_ListByConfigurationStore" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppConfigurationStoreResource created on azure
// for more information of creating AppConfigurationStoreResource, please refer to the document of AppConfigurationStoreResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
string resourceGroupName = "myResourceGroup";
string configStoreName = "contoso";
ResourceIdentifier appConfigurationStoreResourceId = AppConfigurationStoreResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configStoreName);
AppConfigurationStoreResource appConfigurationStore = client.GetAppConfigurationStoreResource(appConfigurationStoreResourceId);

// get the collection of this AppConfigurationPrivateLinkResource
AppConfigurationPrivateLinkResourceCollection collection = appConfigurationStore.GetAppConfigurationPrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (AppConfigurationPrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppConfigurationPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
