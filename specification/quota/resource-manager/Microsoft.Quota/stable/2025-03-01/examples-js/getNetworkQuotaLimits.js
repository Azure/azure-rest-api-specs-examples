const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of current quota limits of all resources for the specified scope. The response from this GET operation can be leveraged to submit requests to update a quota.
 *
 * @summary Get a list of current quota limits of all resources for the specified scope. The response from this GET operation can be leveraged to submit requests to update a quota.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/getNetworkQuotaLimits.json
 */
async function quotasListQuotaLimitsForNetwork() {
  const scope =
    "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.quota.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
