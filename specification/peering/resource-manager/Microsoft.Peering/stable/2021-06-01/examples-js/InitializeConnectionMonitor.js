const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function initializePeeringServiceForConnectionMonitorFunctionality() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.initializeConnectionMonitor();
  console.log(result);
}

initializePeeringServiceForConnectionMonitorFunctionality().catch(console.error);
