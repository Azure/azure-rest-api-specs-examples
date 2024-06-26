const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all connection monitor tests under the given subscription, resource group and peering service.
 *
 * @summary Lists all connection monitor tests under the given subscription, resource group and peering service.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListConnectionMonitorTestsByPeeringService.json
 */
async function listAllConnectionMonitorTestsAssociatedWithThePeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.connectionMonitorTests.listByPeeringService(
    resourceGroupName,
    peeringServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllConnectionMonitorTestsAssociatedWithThePeeringService().catch(console.error);
