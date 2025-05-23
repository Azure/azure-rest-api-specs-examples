const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list DiscoveredAsset resources by subscription ID
 *
 * @summary list DiscoveredAsset resources by subscription ID
 * x-ms-original-file: 2024-09-01-preview/List_DiscoveredAssets_Subscription.json
 */
async function listDiscoveredAssetsSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.discoveredAssets.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
