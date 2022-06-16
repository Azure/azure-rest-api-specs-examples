const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get operation.
 *
 * @summary Get operation.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Operations_Get.json
 */
async function operationsGet() {
  const subscriptionId = "{subscriptionId}";
  const locationName = "{locationName}";
  const name = "{operationName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.operations.get(locationName, name);
  console.log(result);
}

operationsGet().catch(console.error);
