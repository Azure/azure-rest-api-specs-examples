const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager profile.
 *
 * @summary create or update a Traffic Manager profile.
 * x-ms-original-file: 2024-04-01-preview/Profile-PUT-WithNestedEndpoints.json
 */
async function profilePUTWithNestedEndpoints() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate("myresourcegroup", "parentprofile", {
    location: "global",
    dnsConfig: { relativeName: "parentprofile", ttl: 35 },
    endpoints: [
      {
        name: "MyFirstNestedEndpoint",
        type: "Microsoft.Network/trafficManagerProfiles/nestedEndpoints",
        endpointStatus: "Enabled",
        minChildEndpoints: 2,
        minChildEndpointsIPv4: 1,
        minChildEndpointsIPv6: 2,
        priority: 1,
        target: "firstnestedprofile.tmpreview.watmtest.azure-test.net",
        weight: 1,
      },
      {
        name: "MySecondNestedEndpoint",
        type: "Microsoft.Network/trafficManagerProfiles/nestedEndpoints",
        endpointStatus: "Enabled",
        minChildEndpoints: 2,
        minChildEndpointsIPv4: 2,
        minChildEndpointsIPv6: 1,
        priority: 2,
        target: "secondnestedprofile.tmpreview.watmtest.azure-test.net",
        weight: 1,
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
    trafficRoutingMethod: "Priority",
  });
  console.log(result);
}
