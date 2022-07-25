const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the workspaces within a resource group.
 *
 * @summary Gets all the workspaces within a resource group.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspacesListByResourceGroup.json
 */
async function listsWorkspaces() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsWorkspaces().catch(console.error);
