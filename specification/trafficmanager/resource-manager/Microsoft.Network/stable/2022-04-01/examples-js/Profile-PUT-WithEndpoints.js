const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Traffic Manager profile.
 *
 * @summary Create or update a Traffic Manager profile.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PUT-WithEndpoints.json
 */
async function profilePutWithEndpoints() {
  const subscriptionId = process.env["TRAFFICMANAGER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName =
    process.env["TRAFFICMANAGER_RESOURCE_GROUP"] || "azuresdkfornetautoresttrafficmanager2583";
  const profileName = "azuresdkfornetautoresttrafficmanager6192";
  const parameters = {
    dnsConfig: {
      relativeName: "azuresdkfornetautoresttrafficmanager6192",
      ttl: 35,
    },
    endpoints: [
      {
        name: "My external endpoint",
        type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
        endpointLocation: "North Europe",
        endpointStatus: "Enabled",
        target: "foobar.contoso.com",
      },
    ],
    location: "global",
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
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(resourceGroupName, profileName, parameters);
  console.log(result);
}
