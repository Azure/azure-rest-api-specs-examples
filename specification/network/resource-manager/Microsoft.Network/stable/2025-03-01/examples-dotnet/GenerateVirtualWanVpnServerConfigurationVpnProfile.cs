using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/GenerateVirtualWanVpnServerConfigurationVpnProfile.json
// this example is just showing the usage of "Generatevirtualwanvpnserverconfigurationvpnprofile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualWanResource created on azure
// for more information of creating VirtualWanResource, please refer to the document of VirtualWanResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualWanName = "wan1";
ResourceIdentifier virtualWanResourceId = VirtualWanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualWanName);
VirtualWanResource virtualWan = client.GetVirtualWanResource(virtualWanResourceId);

// invoke the operation
VirtualWanVpnProfileContent content = new VirtualWanVpnProfileContent
{
    VpnServerConfigurationResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnconfig1"),
    AuthenticationMethod = NetworkAuthenticationMethod.Eaptls,
};
ArmOperation<VpnProfileResponse> lro = await virtualWan.GenerateVirtualWanVpnServerConfigurationVpnProfileAsync(WaitUntil.Completed, content);
VpnProfileResponse result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
