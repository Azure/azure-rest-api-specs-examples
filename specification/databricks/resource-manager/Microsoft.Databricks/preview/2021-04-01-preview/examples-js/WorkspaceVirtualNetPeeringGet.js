const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the workspace vNet Peering.
 *
 * @summary Gets the workspace vNet Peering.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceVirtualNetPeeringGet.json
 */
async function getAWorkspaceWithVNetPeeringConfigured() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const workspaceName = "myWorkspace";
  const peeringName = "vNetPeering";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.vNetPeering.get(resourceGroupName, workspaceName, peeringName);
  console.log(result);
}

getAWorkspaceWithVNetPeeringConfigured().catch(console.error);
