const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager profile.
 *
 * @summary create or update a Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/Profile-PUT-WithEndpointsAndRecordType.json
 */
async function profilePUTWithEndpointsAndRecordType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(
    "azuresdkfornetautoresttrafficmanager2583",
    "azuresdkfornetautoresttrafficmanager6192",
    {
      profileStatus: "Enabled",
      trafficRoutingMethod: "Performance",
      dnsConfig: { relativeName: "azuresdkfornetautoresttrafficmanager6192", ttl: 35 },
      monitorConfig: {
        protocol: "HTTP",
        port: 80,
        path: "/testpath.aspx",
        intervalInSeconds: 10,
        timeoutInSeconds: 5,
        toleratedNumberOfFailures: 2,
      },
      endpoints: [
        {
          name: "My external endpoint",
          type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
          target: "foobar.contoso.com",
          endpointStatus: "Enabled",
          endpointLocation: "North Europe",
        },
      ],
      recordType: "CNAME",
      location: "global",
    },
  );
  console.log(result);
}
