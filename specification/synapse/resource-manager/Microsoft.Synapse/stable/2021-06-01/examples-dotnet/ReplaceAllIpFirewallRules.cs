using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ReplaceAllIpFirewallRules.json
// this example is just showing the usage of "IpFirewallRules_ReplaceAll" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceResource created on azure
// for more information of creating SynapseWorkspaceResource, please refer to the document of SynapseWorkspaceResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
ResourceIdentifier synapseWorkspaceResourceId = SynapseWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceResource synapseWorkspace = client.GetSynapseWorkspaceResource(synapseWorkspaceResourceId);

// invoke the operation
ReplaceAllIPFirewallRulesContent content = new ReplaceAllIPFirewallRulesContent()
{
    IPFirewallRules =
    {
    ["AnotherExampleFirewallRule"] = new SynapseIPFirewallRuleProperties()
    {
    EndIPAddress = IPAddress.Parse("10.0.1.254"),
    StartIPAddress = IPAddress.Parse("10.0.1.0"),
    },
    ["ExampleFirewallRule"] = new SynapseIPFirewallRuleProperties()
    {
    EndIPAddress = IPAddress.Parse("10.0.0.254"),
    StartIPAddress = IPAddress.Parse("10.0.0.0"),
    },
    },
};
ArmOperation<ReplaceAllFirewallRulesOperationResult> lro = await synapseWorkspace.ReplaceAllIpFirewallRuleAsync(WaitUntil.Completed, content);
ReplaceAllFirewallRulesOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
