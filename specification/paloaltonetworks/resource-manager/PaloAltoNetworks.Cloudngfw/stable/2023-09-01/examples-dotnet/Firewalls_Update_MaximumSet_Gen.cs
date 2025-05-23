using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/Firewalls_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Firewalls_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PaloAltoNetworksFirewallResource created on azure
// for more information of creating PaloAltoNetworksFirewallResource, please refer to the document of PaloAltoNetworksFirewallResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "firewall-rg";
string firewallName = "firewall1";
ResourceIdentifier paloAltoNetworksFirewallResourceId = PaloAltoNetworksFirewallResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallName);
PaloAltoNetworksFirewallResource paloAltoNetworksFirewall = client.GetPaloAltoNetworksFirewallResource(paloAltoNetworksFirewallResourceId);

// invoke the operation
PaloAltoNetworksFirewallPatch patch = new PaloAltoNetworksFirewallPatch
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key16")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["tagName"] = "value"
    },
    Properties = new FirewallUpdateProperties
    {
        PanETag = new ETag("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
        NetworkProfile = new FirewallNetworkProfile(FirewallNetworkType.Vnet, new IPAddressInfo[]
{
new IPAddressInfo
{
ResourceId = new ResourceIdentifier("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1"),
Address = "20.22.92.11",
}
}, AllowEgressNatType.Enabled)
        {
            VnetConfiguration = new FirewallVnetConfiguration(new IPAddressSpaceInfo
            {
                ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet"),
                AddressSpace = "10.1.0.0/16",
            }, new IPAddressSpaceInfo
            {
                ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet"),
                AddressSpace = "10.1.1.0/24",
            }, new IPAddressSpaceInfo
            {
                ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
                AddressSpace = "10.1.1.0/24",
            })
            {
                IPOfTrustSubnetForUdr = new IPAddressInfo
                {
                    ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
                    Address = "10.1.1.0/24",
                },
            },
            VwanConfiguration = new FirewallVwanConfiguration(new IPAddressSpaceInfo
            {
                ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
                AddressSpace = "10.1.1.0/24",
            })
            {
                NetworkVirtualApplianceId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
                TrustSubnet = new IPAddressSpaceInfo
                {
                    ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet"),
                    AddressSpace = "10.1.1.0/24",
                },
                UnTrustSubnet = new IPAddressSpaceInfo
                {
                    ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
                    AddressSpace = "10.1.1.0/24",
                },
                IPOfTrustSubnetForUdr = new IPAddressInfo
                {
                    ResourceId = new ResourceIdentifier("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
                    Address = "10.1.1.0/24",
                },
            },
            EgressNatIP = {new IPAddressInfo
            {
            ResourceId = new ResourceIdentifier("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1"),
            Address = "20.22.92.111",
            }},
            TrustedRanges = { "20.22.92.11" },
        },
        IsPanoramaManaged = FirewallBooleanType.True,
        PanoramaConfig = new FirewallPanoramaConfiguration("bas64EncodedString"),
        AssociatedRulestack = new RulestackDetails
        {
            ResourceId = new ResourceIdentifier("aaaaaaaaaa"),
            RulestackId = "aaaaaaaaaaaaaaaa",
            Location = new AzureLocation("eastus"),
        },
        DnsSettings = new FirewallDnsSettings
        {
            EnableDnsProxy = AllowDnsProxyType.Disabled,
            EnabledDnsType = EnabledDnsType.Custom,
            DnsServers = {new IPAddressInfo
            {
            ResourceId = new ResourceIdentifier("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1"),
            Address = "20.22.92.111",
            }},
        },
        FrontEndSettings = {new FirewallFrontendSetting("frontendsetting11", FirewallProtocolType.Tcp, new FirewallEndpointConfiguration("80", new IPAddressInfo
        {
        ResourceId = new ResourceIdentifier("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp1"),
        Address = "20.22.91.251",
        }), new FirewallEndpointConfiguration("80", new IPAddressInfo
        {
        ResourceId = new ResourceIdentifier("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp2"),
        Address = "20.22.32.136",
        }))},
        PlanData = new FirewallBillingPlanInfo(FirewallBillingCycle.Weekly, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
        {
            UsageType = FirewallBillingPlanUsageType.Payg,
        },
        MarketplaceDetails = new PanFirewallMarketplaceDetails("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa")
        {
            MarketplaceSubscriptionStatus = MarketplaceSubscriptionStatus.PendingFulfillmentStart,
        },
    },
};
PaloAltoNetworksFirewallResource result = await paloAltoNetworksFirewall.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PaloAltoNetworksFirewallData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
