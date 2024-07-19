using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/RouteMapDelete.json
// this example is just showing the usage of "RouteMaps_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RouteMapResource created on azure
// for more information of creating RouteMapResource, please refer to the document of RouteMapResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string routeMapName = "routeMap1";
ResourceIdentifier routeMapResourceId = RouteMapResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, routeMapName);
RouteMapResource routeMap = client.GetRouteMapResource(routeMapResourceId);

// invoke the operation
await routeMap.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
