const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a DDoS protection plan tags.
 *
 * @summary Update a DDoS protection plan tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/DdosProtectionPlanUpdateTags.json
 */
async function dDoSProtectionPlanUpdateTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ddosProtectionPlanName = "test-plan";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosProtectionPlans.updateTags(
    resourceGroupName,
    ddosProtectionPlanName,
    parameters
  );
  console.log(result);
}

dDoSProtectionPlanUpdateTags().catch(console.error);
