const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function nameAvailabilityTestNameNotAvailablePost23() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const parameters = {
    name: "azsmnet4696",
    type: "microsoft.network/trafficmanagerprofiles",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.checkTrafficManagerRelativeDnsNameAvailability(parameters);
  console.log(result);
}

nameAvailabilityTestNameNotAvailablePost23().catch(console.error);
