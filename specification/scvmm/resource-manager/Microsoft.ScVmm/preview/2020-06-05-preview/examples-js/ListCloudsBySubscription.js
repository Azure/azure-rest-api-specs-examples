const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of Clouds in a subscription.
 *
 * @summary List of Clouds in a subscription.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListCloudsBySubscription.json
 */
async function listCloudsBySubscription() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clouds.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
