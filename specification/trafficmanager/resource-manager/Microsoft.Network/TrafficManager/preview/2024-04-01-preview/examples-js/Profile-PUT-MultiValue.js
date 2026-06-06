const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager profile.
 *
 * @summary create or update a Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/Profile-PUT-MultiValue.json
 */
async function profilePUTMultiValue() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(
    "azuresdkfornetautoresttrafficmanager1421",
    "azsmnet6386",
    {
      location: "global",
      dnsConfig: { relativeName: "azsmnet6386", ttl: 35 },
      maxReturn: 2,
      monitorConfig: { path: "/testpath.aspx", port: 80, protocol: "HTTP" },
      profileStatus: "Enabled",
      trafficRoutingMethod: "MultiValue",
      trafficViewEnrollmentStatus: "Disabled",
    },
  );
  console.log(result);
}
