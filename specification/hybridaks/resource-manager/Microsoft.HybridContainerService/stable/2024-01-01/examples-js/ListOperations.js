const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the supported operations
 *
 * @summary Lists the supported operations
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/ListOperations.json
 */
async function listOperations() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
