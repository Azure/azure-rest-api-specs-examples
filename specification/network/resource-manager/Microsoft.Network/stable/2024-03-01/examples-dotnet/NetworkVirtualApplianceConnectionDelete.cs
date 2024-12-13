using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkVirtualApplianceConnectionDelete.json
// this example is just showing the usage of "NetworkVirtualApplianceConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkVirtualApplianceConnectionResource created on azure
// for more information of creating NetworkVirtualApplianceConnectionResource, please refer to the document of NetworkVirtualApplianceConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkVirtualApplianceName = "nva1";
string connectionName = "connection1";
ResourceIdentifier networkVirtualApplianceConnectionResourceId = NetworkVirtualApplianceConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkVirtualApplianceName, connectionName);
NetworkVirtualApplianceConnectionResource networkVirtualApplianceConnection = client.GetNetworkVirtualApplianceConnectionResource(networkVirtualApplianceConnectionResourceId);

// invoke the operation
await networkVirtualApplianceConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
