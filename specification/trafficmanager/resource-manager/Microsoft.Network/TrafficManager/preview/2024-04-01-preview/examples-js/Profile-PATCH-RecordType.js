const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Traffic Manager profile.
 *
 * @summary update a Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/Profile-PATCH-RecordType.json
 */
async function profilePatchRecordType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.update(
    "azuresdkfornetautoresttrafficmanager2583",
    "azuresdkfornetautoresttrafficmanager6192",
    { recordType: "CNAME" },
  );
  console.log(result);
}
