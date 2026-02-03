using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/P2SVpnGatewayGetConnectionHealthDetailed.json
// this example is just showing the usage of "P2sVpnGateways_GetP2SVpnConnectionHealthDetailed" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this P2SVpnGatewayResource created on azure
// for more information of creating P2SVpnGatewayResource, please refer to the document of P2SVpnGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "p2s-vpn-gateway-test";
string gatewayName = "p2svpngateway";
ResourceIdentifier p2sVpnGatewayResourceId = P2SVpnGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName);
P2SVpnGatewayResource p2sVpnGateway = client.GetP2SVpnGatewayResource(p2sVpnGatewayResourceId);

// invoke the operation
P2SVpnConnectionHealthContent content = new P2SVpnConnectionHealthContent
{
    VpnUserNamesFilter = { "vpnUser1", "vpnUser2" },
    OutputBlobSasUri = new Uri("https://blobcortextesturl.blob.core.windows.net/folderforconfig/p2sconnectionhealths?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b"),
};
ArmOperation<P2SVpnConnectionHealth> lro = await p2sVpnGateway.GetP2SVpnConnectionHealthDetailedAsync(WaitUntil.Completed, content);
P2SVpnConnectionHealth result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
