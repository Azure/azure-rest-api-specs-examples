const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the availability of a Traffic Manager Relative DNS name.
 *
 * @summary Checks the availability of a Traffic Manager Relative DNS name.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/NameAvailabilityTest_NameNotAvailable-POST-example-23.json
 */
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
