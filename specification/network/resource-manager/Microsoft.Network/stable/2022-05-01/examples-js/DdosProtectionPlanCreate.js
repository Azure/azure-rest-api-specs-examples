const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a DDoS protection plan.
 *
 * @summary Creates or updates a DDoS protection plan.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/DdosProtectionPlanCreate.json
 */
async function createDDoSProtectionPlan() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ddosProtectionPlanName = "test-plan";
  const parameters = { location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosProtectionPlans.beginCreateOrUpdateAndWait(
    resourceGroupName,
    ddosProtectionPlanName,
    parameters
  );
  console.log(result);
}

createDDoSProtectionPlan().catch(console.error);
