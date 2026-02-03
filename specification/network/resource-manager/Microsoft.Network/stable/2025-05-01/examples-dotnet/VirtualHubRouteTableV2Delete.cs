using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualHubRouteTableV2Delete.json
// this example is just showing the usage of "VirtualHubRouteTableV2s_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualHubRouteTableV2Resource created on azure
// for more information of creating VirtualHubRouteTableV2Resource, please refer to the document of VirtualHubRouteTableV2Resource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string routeTableName = "virtualHubRouteTable1a";
ResourceIdentifier virtualHubRouteTableV2ResourceId = VirtualHubRouteTableV2Resource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, routeTableName);
VirtualHubRouteTableV2Resource virtualHubRouteTableV2 = client.GetVirtualHubRouteTableV2Resource(virtualHubRouteTableV2ResourceId);

// invoke the operation
await virtualHubRouteTableV2.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
