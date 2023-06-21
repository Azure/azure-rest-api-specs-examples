const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the availability of a Traffic Manager Relative DNS name.
 *
 * @summary Checks the availability of a Traffic Manager Relative DNS name.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/NameAvailabilityTest_NameAvailable-POST-example-21.json
 */
async function nameAvailabilityTestNameAvailablePost21() {
  const subscriptionId =
    process.env["TRAFFICMANAGER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const parameters = {
    name: "azsmnet5403",
    type: "microsoft.network/trafficmanagerprofiles",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.checkTrafficManagerRelativeDnsNameAvailability(parameters);
  console.log(result);
}
