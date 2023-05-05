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

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/Firewalls_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "Firewalls_Delete" operation, for the dependent resources, they will have to be created separately.

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
await firewallResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
