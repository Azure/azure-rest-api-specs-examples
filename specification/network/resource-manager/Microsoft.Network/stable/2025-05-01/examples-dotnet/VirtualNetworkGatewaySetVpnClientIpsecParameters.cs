using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkGatewaySetVpnClientIpsecParameters.json
// this example is just showing the usage of "VirtualNetworkGateways_SetVpnclientIPsecParameters" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkGatewayResource created on azure
// for more information of creating VirtualNetworkGatewayResource, please refer to the document of VirtualNetworkGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkGatewayName = "vpngw";
ResourceIdentifier virtualNetworkGatewayResourceId = VirtualNetworkGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkGatewayName);
VirtualNetworkGatewayResource virtualNetworkGateway = client.GetVirtualNetworkGatewayResource(virtualNetworkGatewayResourceId);

// invoke the operation
VpnClientIPsecParameters vpnclientIPsecParams = new VpnClientIPsecParameters(
    86473,
    429497,
    IPsecEncryption.Aes256,
    IPsecIntegrity.Sha256,
    IkeEncryption.Aes256,
    IkeIntegrity.Sha384,
    DHGroup.DHGroup2,
    PfsGroup.Pfs2);
ArmOperation<VpnClientIPsecParameters> lro = await virtualNetworkGateway.SetVpnclientIPsecParametersAsync(WaitUntil.Completed, vpnclientIPsecParams);
VpnClientIPsecParameters result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
