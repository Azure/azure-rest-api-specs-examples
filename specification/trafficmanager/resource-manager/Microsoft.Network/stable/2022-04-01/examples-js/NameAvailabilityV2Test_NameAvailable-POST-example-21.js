const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the availability of a Traffic Manager Relative DNS name.
 *
 * @summary Checks the availability of a Traffic Manager Relative DNS name.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/NameAvailabilityV2Test_NameAvailable-POST-example-21.json
 */
async function nameAvailabilityV2TestNameAvailablePost21() {
  const subscriptionId = process.env["TRAFFICMANAGER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const parameters = {
    name: "azsmnet5403",
    type: "microsoft.network/trafficmanagerprofiles",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.checkTrafficManagerNameAvailabilityV2(parameters);
  console.log(result);
}
