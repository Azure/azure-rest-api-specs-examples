const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the quota and actual usage of the CDN profiles under the given subscription.
 *
 * @summary Check the quota and actual usage of the CDN profiles under the given subscription.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/ResourceUsage_List.json
 */
async function resourceUsageList() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceUsageOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
