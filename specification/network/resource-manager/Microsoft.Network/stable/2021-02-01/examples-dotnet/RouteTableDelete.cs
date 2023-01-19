using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/examples/RouteTableDelete.json
// this example is just showing the usage of "RouteTables_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RouteTableResource created on azure
// for more information of creating RouteTableResource, please refer to the document of RouteTableResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string routeTableName = "testrt";
ResourceIdentifier routeTableResourceId = RouteTableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, routeTableName);
RouteTableResource routeTable = client.GetRouteTableResource(routeTableResourceId);

// invoke the operation
await routeTable.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
