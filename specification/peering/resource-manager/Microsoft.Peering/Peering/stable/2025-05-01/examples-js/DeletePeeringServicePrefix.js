const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes an existing prefix with the specified name under the given subscription, resource group and peering service.
 *
 * @summary deletes an existing prefix with the specified name under the given subscription, resource group and peering service.
 * x-ms-original-file: 2025-05-01/DeletePeeringServicePrefix.json
 */
async function deleteAPrefixAssociatedWithThePeeringService() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new PeeringManagementClient(credential, subscriptionId);
  await client.prefixes.delete("rgName", "peeringServiceName", "peeringServicePrefixName");
}
