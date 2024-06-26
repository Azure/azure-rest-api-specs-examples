const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the workspaces within a resource group.
 *
 * @summary Gets all the workspaces within a resource group.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspacesListByResourceGroup.json
 */
async function listsWorkspaces() {
  const subscriptionId = process.env["DATABRICKS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["DATABRICKS_RESOURCE_GROUP"] || "rg";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
