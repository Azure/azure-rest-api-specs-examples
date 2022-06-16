const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteConnectionMonitorTest() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const connectionMonitorTestName = "connectionMonitorTestName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitorTests.delete(
    resourceGroupName,
    peeringServiceName,
    connectionMonitorTestName
  );
  console.log(result);
}

deleteConnectionMonitorTest().catch(console.error);
