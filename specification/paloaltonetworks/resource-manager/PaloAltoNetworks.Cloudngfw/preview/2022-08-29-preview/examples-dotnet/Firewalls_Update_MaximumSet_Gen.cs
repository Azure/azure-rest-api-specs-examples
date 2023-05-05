using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/Firewalls_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Firewalls_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirewallResource created on azure
// for more information of creating FirewallResource, please refer to the document of FirewallResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "firewall-rg";
string firewallName = "firewall1";
ResourceIdentifier firewallResourceId = FirewallResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallName);
FirewallResource firewallResource = client.GetFirewallResource(firewallResourceId);

// invoke the operation
FirewallResourcePatch patch = new FirewallResourcePatch()
{
    Identity = new AzureResourceManagerManagedIdentityProperties(ManagedIdentityType.None)
    {
        UserAssignedIdentities =
        {
        ["key16"] = new UserAssignedIdentity(),
        },
    },
    Tags =
    {
    ["tagName"] = "value",
    },
    Properties = new FirewallResourceUpdateProperties()
    {
        PanETag = "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
        NetworkProfile = new NetworkProfile(NetworkType.Vnet, new IPAddress[]
{
new IPAddress()
{
ResourceId = "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1",
Address = "20.22.92.11",
}
}, EgressNat.Enabled)
        {
            VnetConfiguration = new VnetConfiguration(new IPAddressSpace()
            {
                ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet",
                AddressSpace = "10.1.0.0/16",
            }, new IPAddressSpace()
            {
                ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet",
                AddressSpace = "10.1.1.0/24",
            }, new IPAddressSpace()
            {
                ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
                AddressSpace = "10.1.1.0/24",
            })
            {
                IPOfTrustSubnetForUdr = new IPAddress()
                {
                    ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
                    Address = "10.1.1.0/24",
                },
            },
            VwanConfiguration = new VwanConfiguration(new IPAddressSpace()
            {
                ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
                AddressSpace = "10.1.1.0/24",
            })
            {
                NetworkVirtualApplianceId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
                TrustSubnet = new IPAddressSpace()
                {
                    ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet",
                    AddressSpace = "10.1.1.0/24",
                },
                UnTrustSubnet = new IPAddressSpace()
                {
                    ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
                    AddressSpace = "10.1.1.0/24",
                },
                IPOfTrustSubnetForUdr = new IPAddress()
                {
                    ResourceId = "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
                    Address = "10.1.1.0/24",
                },
            },
            EgressNatIP =
            {
            new IPAddress()
            {
            ResourceId = "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1",
            Address = "20.22.92.111",
            }
            },
        },
        IsPanoramaManaged = BooleanEnum.True,
        PanoramaConfig = new PanoramaConfig("bas64EncodedString"),
        AssociatedRulestack = new RulestackDetails()
        {
            ResourceId = "aaaaaaaaaa",
            RulestackId = "aaaaaaaaaaaaaaaa",
            Location = new AzureLocation("eastus"),
        },
        DnsSettings = new DnsSettings()
        {
            EnableDnsProxy = DnsProxy.Disabled,
            EnabledDnsType = EnabledDnsType.Custom,
            DnsServers =
            {
            new IPAddress()
            {
            ResourceId = "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1",
            Address = "20.22.92.111",
            }
            },
        },
        FrontEndSettings =
        {
        new FrontendSetting("frontendsetting11",ProtocolType.TCP,new EndpointConfiguration("80",new IPAddress()
        {
        ResourceId = "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp1",
        Address = "20.22.91.251",
        }),new EndpointConfiguration("80",new IPAddress()
        {
        ResourceId = "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp2",
        Address = "20.22.32.136",
        }))
        },
        PlanData = new PlanData(BillingCycle.Weekly, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
        {
            UsageType = UsageType.Payg,
        },
        MarketplaceDetails = new MarketplaceDetails("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa")
        {
            MarketplaceSubscriptionStatus = MarketplaceSubscriptionStatus.PendingFulfillmentStart,
        },
    },
};
FirewallResource result = await firewallResource.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirewallResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
