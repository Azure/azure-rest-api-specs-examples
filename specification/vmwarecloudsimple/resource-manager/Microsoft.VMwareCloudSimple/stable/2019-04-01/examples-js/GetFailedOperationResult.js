const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return an async operation
 *
 * @summary Return an async operation
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetFailedOperationResult.json
 */
async function getFailedOperationResult() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const regionId = "westus2";
  const referer = "https://management.azure.com/";
  const operationId = "d030bb3f-7d53-11e9-8e09-9a86872085ff";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.operations.get(regionId, referer, operationId);
  console.log(result);
}
