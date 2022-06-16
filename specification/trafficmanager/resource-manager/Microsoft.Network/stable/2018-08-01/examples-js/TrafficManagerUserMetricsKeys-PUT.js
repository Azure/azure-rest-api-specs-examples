const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function trafficManagerUserMetricsKeysPut() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.trafficManagerUserMetricsKeys.createOrUpdate();
  console.log(result);
}

trafficManagerUserMetricsKeysPut().catch(console.error);
