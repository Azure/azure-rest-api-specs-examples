const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a DDoS custom policy.
 *
 * @summary Creates or updates a DDoS custom policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DdosCustomPolicyCreate.json
 */
async function createDDoSCustomPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ddosCustomPolicyName = "test-ddos-custom-policy";
  const parameters = {
    location: "centraluseuap",
    protocolCustomSettings: [{ protocol: "Tcp" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosCustomPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    ddosCustomPolicyName,
    parameters
  );
  console.log(result);
}

createDDoSCustomPolicy().catch(console.error);
