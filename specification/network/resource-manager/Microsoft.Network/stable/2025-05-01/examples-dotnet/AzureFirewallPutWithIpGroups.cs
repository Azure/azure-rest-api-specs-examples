using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/AzureFirewallPutWithIpGroups.json
// this example is just showing the usage of "AzureFirewalls_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AzureFirewallResource
AzureFirewallCollection collection = resourceGroupResource.GetAzureFirewalls();

// invoke the operation
string azureFirewallName = "azurefirewall";
AzureFirewallData data = new AzureFirewallData
{
    Zones = { },
    ApplicationRuleCollections = {new AzureFirewallApplicationRuleCollectionData
    {
    Priority = 110,
    ActionType = AzureFirewallRCActionType.Deny,
    Rules = {new AzureFirewallApplicationRule
    {
    Name = "rule1",
    Description = "Deny inbound rule",
    SourceAddresses = {"216.58.216.164", "10.0.0.0/24"},
    Protocols = {new AzureFirewallApplicationRuleProtocol
    {
    ProtocolType = AzureFirewallApplicationRuleProtocolType.Https,
    Port = 443,
    }},
    TargetFqdns = {"www.test.com"},
    }},
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/applicationRuleCollections/apprulecoll"),
    Name = "apprulecoll",
    }},
    NatRuleCollections = {new AzureFirewallNatRuleCollectionData
    {
    Priority = 112,
    ActionType = AzureFirewallNatRCActionType.Dnat,
    Rules = {new AzureFirewallNatRule
    {
    Name = "DNAT-HTTPS-traffic",
    Description = "D-NAT all outbound web traffic for inspection",
    SourceAddresses = {"*"},
    DestinationAddresses = {"1.2.3.4"},
    DestinationPorts = {"443"},
    Protocols = {AzureFirewallNetworkRuleProtocol.Tcp},
    TranslatedAddress = "1.2.3.5",
    TranslatedPort = "8443",
    }, new AzureFirewallNatRule
    {
    Name = "DNAT-HTTP-traffic-With-FQDN",
    Description = "D-NAT all inbound web traffic for inspection",
    SourceAddresses = {"*"},
    DestinationAddresses = {"1.2.3.4"},
    DestinationPorts = {"80"},
    Protocols = {AzureFirewallNetworkRuleProtocol.Tcp},
    TranslatedPort = "880",
    TranslatedFqdn = "internalhttpserver",
    }},
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/natRuleCollections/natrulecoll"),
    Name = "natrulecoll",
    }},
    NetworkRuleCollections = {new AzureFirewallNetworkRuleCollectionData
    {
    Priority = 112,
    ActionType = AzureFirewallRCActionType.Deny,
    Rules = {new AzureFirewallNetworkRule
    {
    Name = "L4-traffic",
    Description = "Block traffic based on source IPs and ports",
    Protocols = {AzureFirewallNetworkRuleProtocol.Tcp},
    SourceAddresses = {"192.168.1.1-192.168.1.12", "10.1.4.12-10.1.4.255"},
    DestinationAddresses = {"*"},
    DestinationPorts = {"443-444", "8443"},
    }, new AzureFirewallNetworkRule
    {
    Name = "L4-traffic-with-FQDN",
    Description = "Block traffic based on source IPs and ports to amazon",
    Protocols = {AzureFirewallNetworkRuleProtocol.Tcp},
    SourceAddresses = {"10.2.4.12-10.2.4.255"},
    DestinationPorts = {"443-444", "8443"},
    DestinationFqdns = {"www.amazon.com"},
    }},
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/networkRuleCollections/netrulecoll"),
    Name = "netrulecoll",
    }},
    IPConfigurations = {new AzureFirewallIPConfiguration
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"),
    PublicIPAddressId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
    Name = "azureFirewallIpConfiguration",
    }},
    ThreatIntelMode = AzureFirewallThreatIntelMode.Alert,
    Sku = new AzureFirewallSku
    {
        Name = AzureFirewallSkuName.AzfwVnet,
        Tier = AzureFirewallSkuTier.Standard,
    },
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<AzureFirewallResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, azureFirewallName, data);
AzureFirewallResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AzureFirewallData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
