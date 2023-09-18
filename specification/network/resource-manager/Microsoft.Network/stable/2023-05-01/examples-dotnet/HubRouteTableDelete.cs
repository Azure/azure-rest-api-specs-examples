using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/HubRouteTableDelete.json
// this example is just showing the usage of "HubRouteTables_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubRouteTableResource created on azure
// for more information of creating HubRouteTableResource, please refer to the document of HubRouteTableResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string routeTableName = "hubRouteTable1";
ResourceIdentifier hubRouteTableResourceId = HubRouteTableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, routeTableName);
HubRouteTableResource hubRouteTable = client.GetHubRouteTableResource(hubRouteTableResourceId);

// invoke the operation
await hubRouteTable.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
