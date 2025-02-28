const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of current usage for all resources for the scope specified.
 *
 * @summary Get a list of current usage for all resources for the scope specified.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/getComputeUsages.json
 */
async function quotasListUsagesForCompute() {
  const scope =
    "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.usages.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
