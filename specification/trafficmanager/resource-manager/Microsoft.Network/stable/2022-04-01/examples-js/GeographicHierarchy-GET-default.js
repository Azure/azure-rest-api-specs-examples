const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the default Geographic Hierarchy used by the Geographic traffic routing method.
 *
 * @summary Gets the default Geographic Hierarchy used by the Geographic traffic routing method.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/GeographicHierarchy-GET-default.json
 */
async function geographicHierarchyGetDefault() {
  const subscriptionId =
    process.env["TRAFFICMANAGER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.geographicHierarchies.getDefault();
  console.log(result);
}
