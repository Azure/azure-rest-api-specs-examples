const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags of a service endpoint policy.
 *
 * @summary Updates tags of a service endpoint policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ServiceEndpointPolicyUpdateTags.json
 */
async function updateServiceEndpointPolicyTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceEndpointPolicyName = "testServiceEndpointPolicy";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.serviceEndpointPolicies.updateTags(
    resourceGroupName,
    serviceEndpointPolicyName,
    parameters
  );
  console.log(result);
}

updateServiceEndpointPolicyTags().catch(console.error);
