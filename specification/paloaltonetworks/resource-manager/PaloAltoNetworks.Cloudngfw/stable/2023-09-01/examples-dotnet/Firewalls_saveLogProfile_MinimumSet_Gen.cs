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

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/Firewalls_saveLogProfile_MinimumSet_Gen.json
// this example is just showing the usage of "Firewalls_saveLogProfile" operation, for the dependent resources, they will have to be created separately.

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
await paloAltoNetworksFirewall.SaveLogProfileAsync();

Console.WriteLine($"Succeeded");
