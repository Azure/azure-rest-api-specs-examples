using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2025-10-08/examples/MetricsObjectFirewall_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "MetricsObjectFirewall_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MetricsObjectFirewallResource created on azure
// for more information of creating MetricsObjectFirewallResource, please refer to the document of MetricsObjectFirewallResource
string subscriptionId = "aaaaaaa";
string resourceGroupName = "rgopenapi";
string firewallName = "aaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier metricsObjectFirewallResourceId = MetricsObjectFirewallResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallName);
MetricsObjectFirewallResource metricsObjectFirewall = client.GetMetricsObjectFirewallResource(metricsObjectFirewallResourceId);

// invoke the operation
await metricsObjectFirewall.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
