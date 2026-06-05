const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager profile.
 *
 * @summary create or update a Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/Profile-PUT-WithCustomHeaders.json
 */
async function profilePUTWithCustomHeaders() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(
    "azuresdkfornetautoresttrafficmanager2583",
    "azuresdkfornetautoresttrafficmanager6192",
    {
      location: "global",
      dnsConfig: { relativeName: "azuresdkfornetautoresttrafficmanager6192", ttl: 35 },
      endpoints: [
        {
          name: "My external endpoint",
          type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
          customHeaders: [{ name: "header-2", value: "value-2-overridden" }],
          endpointLocation: "North Europe",
          endpointStatus: "Enabled",
          target: "foobar.contoso.com",
        },
      ],
      monitorConfig: {
        path: "/testpath.aspx",
        customHeaders: [
          { name: "header-1", value: "value-1" },
          { name: "header-2", value: "value-2" },
        ],
        expectedStatusCodeRanges: [
          { max: 205, min: 200 },
          { max: 410, min: 400 },
        ],
        intervalInSeconds: 10,
        port: 80,
        timeoutInSeconds: 5,
        toleratedNumberOfFailures: 2,
        protocol: "HTTP",
      },
      profileStatus: "Enabled",
      trafficRoutingMethod: "Performance",
      trafficViewEnrollmentStatus: "Disabled",
    },
  );
  console.log(result);
}
