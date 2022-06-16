const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing prefix with the specified name under the given subscription, resource group and peering service.
 *
 * @summary Gets an existing prefix with the specified name under the given subscription, resource group and peering service.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/GetPeeringServicePrefix.json
 */
async function getAPrefixAssociatedWithThePeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const prefixName = "peeringServicePrefixName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.prefixes.get(resourceGroupName, peeringServiceName, prefixName);
  console.log(result);
}

getAPrefixAssociatedWithThePeeringService().catch(console.error);
