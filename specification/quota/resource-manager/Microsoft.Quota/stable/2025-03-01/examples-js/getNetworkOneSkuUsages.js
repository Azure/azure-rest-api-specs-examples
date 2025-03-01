const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the current usage of a resource.
 *
 * @summary Get the current usage of a resource.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/getNetworkOneSkuUsages.json
 */
async function quotasUsagesRequestForNetwork() {
  const resourceName = "MinPublicIpInterNetworkPrefixLength";
  const scope =
    "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const result = await client.usages.get(resourceName, scope);
  console.log(result);
}
