const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an express route circuit.
 *
 * @summary creates or updates an express route circuit.
 * x-ms-original-file: 2025-09-01/ExpressRouteMultiCloudCircuitCreateWithActivationKey.json
 */
async function createMultiCloudExpressRouteCircuitWithActivationKey() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.createOrUpdate("rg1", "circuitName", {
    serviceProviderProperties: {
      serviceProviderName: "AWS",
      peeringLocation: "uswest2",
      bandwidthInMbps: 200,
    },
    activationKey:
      "eyJzaGFyZWRDb25uZWN0aW9uVXVpZCI6IjE1ODliMDhhLTNmYWQtNDkzNi05MGQyLWE5ZDg3Y2JkNmM3MCIsImNvbm5lY3Rpb25TaXplTWJwcyI6MTAwMC4wLCJkZXN0aW5hdGlvbkFjY291bnRJZCI6IjEyMzQ1Njc4OSIsImVudmlyb25tZW50IjoidXN3ZXN0MiIsImRlc3RpbmF0aW9uRW52aXJvbm1lbnRVcmkiOiIvc3Vic2NyaXB0aW9ucy85OWMzMzc3Ni05ZjRlLTRlNTgtYWJlOC05MjYzZGIxYjljNmUvcmVzb3VyY2VHcm91cHMvQ3Jvc3NDb25uZWN0aW9uLXVzd2VzdDIvcHJvdmlkZXJzL01pY3Jvc29mdC5OZXR3b3JrL2V4cHJlc3NSb3V0ZUNyb3NzQ29ubmVjdGlvbnMvMTU4OWIwOGEtM2ZhZC00OTM2LTkwZDItYTlkODdjYmQ2YzcwIiwidmVyc2lvbiI6MX0=",
    sku: { name: "MultiCloud_MeteredData", tier: "MultiCloud", family: "MeteredData" },
    location: "eastus2euap",
  });
  console.log(result);
}
