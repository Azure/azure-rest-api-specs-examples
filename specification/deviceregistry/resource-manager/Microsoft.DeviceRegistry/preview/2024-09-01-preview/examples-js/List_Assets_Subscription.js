const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list Asset resources by subscription ID
 *
 * @summary list Asset resources by subscription ID
 * x-ms-original-file: 2024-09-01-preview/List_Assets_Subscription.json
 */
async function listAssetsSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assets.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
