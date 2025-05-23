const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the specified service endpoint policy.
 *
 * @summary Deletes the specified service endpoint policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ServiceEndpointPolicyDelete.json
 */
async function deleteServiceEndpointPolicy() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const serviceEndpointPolicyName = "serviceEndpointPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.serviceEndpointPolicies.beginDeleteAndWait(
    resourceGroupName,
    serviceEndpointPolicyName,
  );
  console.log(result);
}
