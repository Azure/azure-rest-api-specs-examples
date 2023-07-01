using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/HubVirtualNetworkConnectionGet.json
// this example is just showing the usage of "HubVirtualNetworkConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubVirtualNetworkConnectionResource created on azure
// for more information of creating HubVirtualNetworkConnectionResource, please refer to the document of HubVirtualNetworkConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string connectionName = "connection1";
ResourceIdentifier hubVirtualNetworkConnectionResourceId = HubVirtualNetworkConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, connectionName);
HubVirtualNetworkConnectionResource hubVirtualNetworkConnection = client.GetHubVirtualNetworkConnectionResource(hubVirtualNetworkConnectionResourceId);

// invoke the operation
HubVirtualNetworkConnectionResource result = await hubVirtualNetworkConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HubVirtualNetworkConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
