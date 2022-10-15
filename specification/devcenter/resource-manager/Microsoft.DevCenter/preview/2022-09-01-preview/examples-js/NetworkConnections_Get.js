const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a network connection resource
 *
 * @summary Gets a network connection resource
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/NetworkConnections_Get.json
 */
async function networkConnectionsGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const networkConnectionName = "uswest3network";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.networkConnections.get(resourceGroupName, networkConnectionName);
  console.log(result);
}

networkConnectionsGet().catch(console.error);
