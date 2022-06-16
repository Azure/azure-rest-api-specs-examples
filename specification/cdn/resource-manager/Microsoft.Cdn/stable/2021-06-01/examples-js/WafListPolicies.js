const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the protection policies within a resource group.
 *
 * @summary Lists all of the protection policies within a resource group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafListPolicies.json
 */
async function listPoliciesInAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policies.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPoliciesInAResourceGroup().catch(console.error);
