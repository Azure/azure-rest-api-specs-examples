const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets information about the specified DDoS protection plan.
 *
 * @summary Gets information about the specified DDoS protection plan.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/DdosProtectionPlanGet.json
 */
async function getDDoSProtectionPlan() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const ddosProtectionPlanName = "test-plan";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosProtectionPlans.get(resourceGroupName, ddosProtectionPlanName);
  console.log(result);
}
