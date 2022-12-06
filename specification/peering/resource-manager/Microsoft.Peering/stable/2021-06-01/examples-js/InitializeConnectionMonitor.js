const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Initialize Peering Service for Connection Monitor functionality
 *
 * @summary Initialize Peering Service for Connection Monitor functionality
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/InitializeConnectionMonitor.json
 */
async function initializePeeringServiceForConnectionMonitorFunctionality() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.initializeConnectionMonitor();
  console.log(result);
}

initializePeeringServiceForConnectionMonitorFunctionality().catch(console.error);
