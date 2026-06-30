const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve counts for Dsc Nodes.
 *
 * @summary retrieve counts for Dsc Nodes.
 * x-ms-original-file: 2024-10-23/listPagedDscNodeConfigurationCounts.json
 */
async function getNodeNodeConfigurationCounts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.nodeCountInformation.get(
    "rg",
    "myAutomationAccount33",
    "nodeconfiguration",
  );
  console.log(result);
}
