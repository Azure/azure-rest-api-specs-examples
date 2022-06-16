const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Extensions under ArcSetting resource.
 *
 * @summary List all Extensions under ArcSetting resource.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/ListExtensionsByArcSetting.json
 */
async function listExtensionsUnderArcSettingResource() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.extensions.listByArcSetting(
    resourceGroupName,
    clusterName,
    arcSettingName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExtensionsUnderArcSettingResource().catch(console.error);
