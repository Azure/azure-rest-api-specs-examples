using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/Firewalls_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "Firewalls_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "firewall-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PaloAltoNetworksFirewallResource
PaloAltoNetworksFirewallCollection collection = resourceGroupResource.GetPaloAltoNetworksFirewalls();

// invoke the operation
string firewallName = "firewall1";
PaloAltoNetworksFirewallData data = new PaloAltoNetworksFirewallData(new AzureLocation("eastus"), new FirewallNetworkProfile(FirewallNetworkType.Vnet, new IPAddressInfo[]
{
new IPAddressInfo
{
ResourceId = new ResourceIdentifier("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1"),
Address = "20.22.92.11",
}
}, AllowEgressNatType.Enabled), new FirewallDnsSettings(), new FirewallBillingPlanInfo(FirewallBillingCycle.Monthly, "liftrpantestplan"), new PanFirewallMarketplaceDetails("liftr-pan-ame-test", "isvtestuklegacy"));
ArmOperation<PaloAltoNetworksFirewallResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, firewallName, data);
PaloAltoNetworksFirewallResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PaloAltoNetworksFirewallData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
