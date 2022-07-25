const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the workspace vNetPeering.
 *
 * @summary Deletes the workspace vNetPeering.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceVirtualNetworkPeeringDelete.json
 */
async function deleteAWorkspaceVNetPeering() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const workspaceName = "myWorkspace";
  const peeringName = "vNetPeering";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.vNetPeering.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    peeringName
  );
  console.log(result);
}

deleteAWorkspaceVNetPeering().catch(console.error);
