const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing prefix with the specified name under the given subscription, resource group and peering service.
 *
 * @summary Deletes an existing prefix with the specified name under the given subscription, resource group and peering service.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/DeletePeeringServicePrefix.json
 */
async function deleteAPrefixAssociatedWithThePeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const prefixName = "peeringServicePrefixName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.prefixes.delete(resourceGroupName, peeringServiceName, prefixName);
  console.log(result);
}

deleteAPrefixAssociatedWithThePeeringService().catch(console.error);
