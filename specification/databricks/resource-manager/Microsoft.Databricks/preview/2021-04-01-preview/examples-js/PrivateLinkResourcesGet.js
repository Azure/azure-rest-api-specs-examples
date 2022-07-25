const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified private link resource for the given group id (sub-resource)
 *
 * @summary Get the specified private link resource for the given group id (sub-resource)
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/PrivateLinkResourcesGet.json
 */
async function getAPrivateLinkResource() {
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "myWorkspace";
  const groupId = "databricks_ui_api";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, workspaceName, groupId);
  console.log(result);
}

getAPrivateLinkResource().catch(console.error);
