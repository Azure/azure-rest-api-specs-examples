const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all DDoS protection plans in a subscription.
 *
 * @summary Gets all DDoS protection plans in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/DdosProtectionPlanListAll.json
 */
async function listAllDDoSProtectionPlans() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.ddosProtectionPlans.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllDDoSProtectionPlans().catch(console.error);
