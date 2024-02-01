using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppConfiguration;
using Azure.ResourceManager.AppConfiguration.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresListKeys.json
// this example is just showing the usage of "ConfigurationStores_ListKeys" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (AppConfigurationStoreApiKey item in appConfigurationStore.GetKeysAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
