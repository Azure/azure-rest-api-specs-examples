const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the availability of a Traffic Manager Relative DNS name.
 *
 * @summary Checks the availability of a Traffic Manager Relative DNS name.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/NameAvailabilityV2Test_NameNotAvailable-POST-example-23.json
 */
async function nameAvailabilityV2TestNameNotAvailablePost23() {
  const subscriptionId = process.env["TRAFFICMANAGER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const parameters = {
    name: "azsmnet4696",
    type: "microsoft.network/trafficmanagerprofiles",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.checkTrafficManagerNameAvailabilityV2(parameters);
  console.log(result);
}
