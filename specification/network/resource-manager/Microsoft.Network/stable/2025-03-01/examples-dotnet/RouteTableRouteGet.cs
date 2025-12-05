using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/RouteTableRouteGet.json
// this example is just showing the usage of "Routes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RouteResource created on azure
// for more information of creating RouteResource, please refer to the document of RouteResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string routeTableName = "testrt";
string routeName = "route1";
ResourceIdentifier routeResourceId = RouteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, routeTableName, routeName);
RouteResource route = client.GetRouteResource(routeResourceId);

// invoke the operation
RouteResource result = await route.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RouteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
