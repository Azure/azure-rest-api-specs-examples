const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a service endpoint policy definition in the specified service endpoint policy.
 *
 * @summary Creates or updates a service endpoint policy definition in the specified service endpoint policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ServiceEndpointPolicyDefinitionCreate.json
 */
async function createServiceEndpointPolicyDefinition() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const serviceEndpointPolicyName = "testPolicy";
  const serviceEndpointPolicyDefinitionName = "testDefinition";
  const serviceEndpointPolicyDefinitions = {
    description: "Storage Service EndpointPolicy Definition",
    service: "Microsoft.Storage",
    serviceResources: [
      "/subscriptions/subid1",
      "/subscriptions/subid1/resourceGroups/storageRg",
      "/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.serviceEndpointPolicyDefinitions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceEndpointPolicyName,
    serviceEndpointPolicyDefinitionName,
    serviceEndpointPolicyDefinitions,
  );
  console.log(result);
}
