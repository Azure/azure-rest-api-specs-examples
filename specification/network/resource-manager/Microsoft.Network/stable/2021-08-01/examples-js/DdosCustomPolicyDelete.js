const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified DDoS custom policy.
 *
 * @summary Deletes the specified DDoS custom policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/DdosCustomPolicyDelete.json
 */
async function deleteDDoSCustomPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ddosCustomPolicyName = "test-ddos-custom-policy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosCustomPolicies.beginDeleteAndWait(
    resourceGroupName,
    ddosCustomPolicyName
  );
  console.log(result);
}

deleteDDoSCustomPolicy().catch(console.error);
