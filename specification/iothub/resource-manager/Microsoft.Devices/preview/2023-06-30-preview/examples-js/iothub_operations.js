const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available IoT Hub REST API operations.
 *
 * @summary Lists all of the available IoT Hub REST API operations.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_operations.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
