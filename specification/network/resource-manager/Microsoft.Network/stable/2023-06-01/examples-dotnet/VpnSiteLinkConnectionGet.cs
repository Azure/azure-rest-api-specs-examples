using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/VpnSiteLinkConnectionGet.json
// this example is just showing the usage of "VpnSiteLinkConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnSiteLinkConnectionResource created on azure
// for more information of creating VpnSiteLinkConnectionResource, please refer to the document of VpnSiteLinkConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string gatewayName = "gateway1";
string connectionName = "vpnConnection1";
string linkConnectionName = "Connection-Link1";
ResourceIdentifier vpnSiteLinkConnectionResourceId = VpnSiteLinkConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName, connectionName, linkConnectionName);
VpnSiteLinkConnectionResource vpnSiteLinkConnection = client.GetVpnSiteLinkConnectionResource(vpnSiteLinkConnectionResourceId);

// invoke the operation
VpnSiteLinkConnectionResource result = await vpnSiteLinkConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VpnSiteLinkConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
