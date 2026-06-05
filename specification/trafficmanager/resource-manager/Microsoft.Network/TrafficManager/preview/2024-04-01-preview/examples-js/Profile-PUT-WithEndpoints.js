const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager profile.
 *
 * @summary create or update a Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/Profile-PUT-WithEndpoints.json
 */
async function profilePUTWithEndpoints() {
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
          endpointLocation: "North Europe",
          endpointStatus: "Enabled",
          target: "foobar.contoso.com",
        },
      ],
      monitorConfig: {
        path: "/testpath.aspx",
        intervalInSeconds: 10,
        port: 80,
        timeoutInSeconds: 5,
        toleratedNumberOfFailures: 2,
        protocol: "HTTP",
      },
      profileStatus: "Enabled",
      trafficRoutingMethod: "Performance",
    },
  );
  console.log(result);
}
