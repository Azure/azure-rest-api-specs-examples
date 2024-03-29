const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Replaces firewall rules
 *
 * @summary Replaces firewall rules
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ReplaceAllIpFirewallRules.json
 */
async function replaceAllIPFirewallRulesInAWorkspace() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const request = {
    ipFirewallRules: {
      anotherExampleFirewallRule: {
        endIpAddress: "10.0.1.254",
        startIpAddress: "10.0.1.0",
      },
      exampleFirewallRule: {
        endIpAddress: "10.0.0.254",
        startIpAddress: "10.0.0.0",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.ipFirewallRules.beginReplaceAllAndWait(
    resourceGroupName,
    workspaceName,
    request
  );
  console.log(result);
}
