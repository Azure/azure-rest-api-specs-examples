using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VpnSitePut.json
// this example is just showing the usage of "VpnSites_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VpnSiteResource
VpnSiteCollection collection = resourceGroupResource.GetVpnSites();

// invoke the operation
string vpnSiteName = "vpnSite1";
VpnSiteData data = new VpnSiteData
{
    VirtualWanId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
    AddressSpace = new VirtualNetworkAddressSpace
    {
        AddressPrefixes = { "10.0.0.0/16" },
    },
    IsSecuritySite = false,
    VpnSiteLinks = {new VpnSiteLinkData
    {
    LinkProperties = new VpnLinkProviderProperties
    {
    LinkProviderName = "vendor1",
    LinkSpeedInMbps = 0,
    },
    IPAddress = "50.50.50.56",
    Fqdn = "link1.vpnsite1.contoso.com",
    BgpProperties = new VpnLinkBgpSettings
    {
    Asn = 1234L,
    BgpPeeringAddress = "192.168.0.0",
    },
    Name = "vpnSiteLink1",
    }},
    O365BreakOutCategories = new O365BreakOutCategoryPolicies
    {
        Allow = true,
        Optimize = true,
        Default = false,
    },
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<VpnSiteResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vpnSiteName, data);
VpnSiteResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VpnSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
