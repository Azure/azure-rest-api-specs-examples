using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppConfiguration;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-05-01/examples/PrivateLinkResourceGet.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppConfigurationPrivateLinkResource created on azure
// for more information of creating AppConfigurationPrivateLinkResource, please refer to the document of AppConfigurationPrivateLinkResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
string resourceGroupName = "myResourceGroup";
string configStoreName = "contoso";
string groupName = "configurationStores";
ResourceIdentifier appConfigurationPrivateLinkResourceId = AppConfigurationPrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configStoreName, groupName);
AppConfigurationPrivateLinkResource appConfigurationPrivateLinkResource = client.GetAppConfigurationPrivateLinkResource(appConfigurationPrivateLinkResourceId);

// invoke the operation
AppConfigurationPrivateLinkResource result = await appConfigurationPrivateLinkResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppConfigurationPrivateLinkResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
