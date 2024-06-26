const { PlaywrightTestingClient } = require("@azure/arm-playwrighttesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get quota by name.
 *
 * @summary Get quota by name.
 * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Quotas_Get.json
 */
async function quotasGet() {
  const subscriptionId =
    process.env["PLAYWRIGHTTESTING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const name = "ScalableExecution";
  const credential = new DefaultAzureCredential();
  const client = new PlaywrightTestingClient(credential, subscriptionId);
  const result = await client.quotas.get(location, name);
  console.log(result);
}
