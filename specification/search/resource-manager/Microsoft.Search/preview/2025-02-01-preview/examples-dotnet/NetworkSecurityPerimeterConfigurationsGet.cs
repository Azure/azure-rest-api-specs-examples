using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Search;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/NetworkSecurityPerimeterConfigurationsGet.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating NetworkSecurityPerimeterConfigurationResource, please refer to the document of NetworkSecurityPerimeterConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string searchServiceName = "mysearchservice";
string nspConfigName = "00000001-2222-3333-4444-111144444444.assoc1";
ResourceIdentifier networkSecurityPerimeterConfigurationResourceId = NetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, searchServiceName, nspConfigName);
NetworkSecurityPerimeterConfigurationResource networkSecurityPerimeterConfiguration = client.GetNetworkSecurityPerimeterConfigurationResource(networkSecurityPerimeterConfigurationResourceId);

// invoke the operation
NetworkSecurityPerimeterConfigurationResource result = await networkSecurityPerimeterConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkSecurityPerimeterConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
