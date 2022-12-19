const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a firewall rule
 *
 * @summary Creates or updates a firewall rule
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateIpFirewallRule.json
 */
async function createAnIPFirewallRule() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const ruleName = "ExampleIpFirewallRule";
  const ipFirewallRuleInfo = {
    endIpAddress: "10.0.0.254",
    startIpAddress: "10.0.0.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.ipFirewallRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    ruleName,
    ipFirewallRuleInfo
  );
  console.log(result);
}
