const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available managed rule sets.
 *
 * @summary Lists all available managed rule sets.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/WafListManagedRuleSets.json
 */
async function listPoliciesInAResourceGroup() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedRuleSets.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
