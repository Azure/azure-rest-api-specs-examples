using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NetworkVirtualApplianceConnectionGet.json
// this example is just showing the usage of "NetworkVirtualApplianceConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkVirtualApplianceResource created on azure
// for more information of creating NetworkVirtualApplianceResource, please refer to the document of NetworkVirtualApplianceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkVirtualApplianceName = "nva1";
ResourceIdentifier networkVirtualApplianceResourceId = NetworkVirtualApplianceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkVirtualApplianceName);
NetworkVirtualApplianceResource networkVirtualAppliance = client.GetNetworkVirtualApplianceResource(networkVirtualApplianceResourceId);

// get the collection of this NetworkVirtualApplianceConnectionResource
NetworkVirtualApplianceConnectionCollection collection = networkVirtualAppliance.GetNetworkVirtualApplianceConnections();

// invoke the operation
string connectionName = "connection1";
bool result = await collection.ExistsAsync(connectionName);

Console.WriteLine($"Succeeded: {result}");
