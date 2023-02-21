using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteIpFirewallRule.json
// this example is just showing the usage of "IpFirewallRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseIPFirewallRuleInfoResource created on azure
// for more information of creating SynapseIPFirewallRuleInfoResource, please refer to the document of SynapseIPFirewallRuleInfoResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
string ruleName = "ExampleIpFirewallRule";
ResourceIdentifier synapseIPFirewallRuleInfoResourceId = SynapseIPFirewallRuleInfoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, ruleName);
SynapseIPFirewallRuleInfoResource synapseIPFirewallRuleInfo = client.GetSynapseIPFirewallRuleInfoResource(synapseIPFirewallRuleInfoResourceId);

// invoke the operation
ArmOperation<BinaryData> lro = await synapseIPFirewallRuleInfo.DeleteAsync(WaitUntil.Completed);
BinaryData result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
