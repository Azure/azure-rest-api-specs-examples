const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private link resources for a given workspace
 *
 * @summary List private link resources for a given workspace
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/ListPrivateLinkResources.json
 */
async function listPrivateLinkResources() {
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPrivateLinkResources().catch(console.error);
