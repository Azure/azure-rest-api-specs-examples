using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/HubRouteTableGet.json
// this example is just showing the usage of "HubRouteTables_Get" operation, for the dependent resources, they will have to be created separately.

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
HubRouteTableResource result = await hubRouteTable.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HubRouteTableData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
