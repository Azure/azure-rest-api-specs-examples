const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private endpoint connections of the workspace
 *
 * @summary List private endpoint connections of the workspace
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/ListPrivateEndpointConnections.json
 */
async function listPrivateEndpointConnections() {
  const subscriptionId =
    process.env["DATABRICKS_SUBSCRIPTION_ID"] || "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = process.env["DATABRICKS_RESOURCE_GROUP"] || "myResourceGroup";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
