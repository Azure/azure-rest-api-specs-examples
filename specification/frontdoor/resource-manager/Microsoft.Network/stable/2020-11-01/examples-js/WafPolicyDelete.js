const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes Policy
 *
 * @summary Deletes Policy
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafPolicyDelete.json
 */
async function deleteProtectionPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const policyName = "Policy1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.policies.beginDeleteAndWait(resourceGroupName, policyName);
  console.log(result);
}

deleteProtectionPolicy().catch(console.error);
