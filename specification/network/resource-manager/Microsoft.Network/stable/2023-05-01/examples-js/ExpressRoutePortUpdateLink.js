const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified ExpressRoutePort resource.
 *
 * @summary Creates or updates the specified ExpressRoutePort resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortUpdateLink.json
 */
async function expressRoutePortUpdateLink() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const expressRoutePortName = "portName";
  const parameters = {
    bandwidthInGbps: 100,
    billingType: "UnlimitedData",
    encapsulation: "QinQ",
    links: [{ name: "link1", adminState: "Enabled" }],
    location: "westus",
    peeringLocation: "peeringLocationName",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRoutePorts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    expressRoutePortName,
    parameters
  );
  console.log(result);
}
