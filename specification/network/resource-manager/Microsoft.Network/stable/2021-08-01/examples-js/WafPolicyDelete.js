const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes Policy.
 *
 * @summary Deletes Policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/WafPolicyDelete.json
 */
async function deletesAWafPolicyWithinAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const policyName = "Policy1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.webApplicationFirewallPolicies.beginDeleteAndWait(
    resourceGroupName,
    policyName
  );
  console.log(result);
}

deletesAWafPolicyWithinAResourceGroup().catch(console.error);
