const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Microsoft Portal operations API.
 *
 * @summary The Microsoft Portal operations API.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/operationsList.json
 */
async function listThePortalOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listThePortalOperations().catch(console.error);
