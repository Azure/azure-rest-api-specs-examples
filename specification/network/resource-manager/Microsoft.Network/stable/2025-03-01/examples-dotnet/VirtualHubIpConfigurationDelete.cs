using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualHubIpConfigurationDelete.json
// this example is just showing the usage of "VirtualHubIpConfiguration_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubIPConfigurationResource created on azure
// for more information of creating HubIPConfigurationResource, please refer to the document of HubIPConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "hub1";
string ipConfigName = "ipconfig1";
ResourceIdentifier hubIPConfigurationResourceId = HubIPConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, ipConfigName);
HubIPConfigurationResource hubIPConfiguration = client.GetHubIPConfigurationResource(hubIPConfigurationResourceId);

// invoke the operation
await hubIPConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
