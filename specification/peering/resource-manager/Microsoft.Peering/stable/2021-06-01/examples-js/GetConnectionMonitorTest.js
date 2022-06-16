const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
