const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Resources by Resource Group
 *
 * @summary List Resources by Resource Group
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/GetByResourceGroup.json
 */
async function listObjectAnchorsAccountsByResourceGroup() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.objectAnchorsAccounts.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listObjectAnchorsAccountsByResourceGroup().catch(console.error);
