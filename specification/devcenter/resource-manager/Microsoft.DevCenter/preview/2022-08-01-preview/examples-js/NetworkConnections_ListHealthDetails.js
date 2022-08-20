const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists health check status details
 *
 * @summary Lists health check status details
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/NetworkConnections_ListHealthDetails.json
 */
async function networkConnectionsListHealthDetails() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const networkConnectionName = "uswest3network";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkConnections.listHealthDetails(
    resourceGroupName,
    networkConnectionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

networkConnectionsListHealthDetails().catch(console.error);
