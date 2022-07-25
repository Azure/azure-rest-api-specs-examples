const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Kusto pools
 *
 * @summary List all Kusto pools
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListByWorkspace.json
 */
async function listKustoPoolsInAWorkspace() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const workspaceName = "kustorptest";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPools.listByWorkspace(resourceGroupName, workspaceName);
  console.log(result);
}

listKustoPoolsInAWorkspace().catch(console.error);
