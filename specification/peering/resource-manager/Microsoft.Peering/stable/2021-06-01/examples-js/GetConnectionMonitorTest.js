const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing connection monitor test with the specified name under the given subscription, resource group and peering service.
 *
 * @summary Gets an existing connection monitor test with the specified name under the given subscription, resource group and peering service.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/GetConnectionMonitorTest.json
 */
async function getConnectionMonitorTest() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const connectionMonitorTestName = "connectionMonitorTestName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitorTests.get(
    resourceGroupName,
    peeringServiceName,
    connectionMonitorTestName
  );
  console.log(result);
}

getConnectionMonitorTest().catch(console.error);
