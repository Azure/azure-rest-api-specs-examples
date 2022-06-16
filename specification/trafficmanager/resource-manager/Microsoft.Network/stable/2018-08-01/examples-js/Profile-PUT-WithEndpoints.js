const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function profilePutWithEndpoints() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager2583";
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

profilePutWithEndpoints().catch(console.error);
