const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a TrafficController
 *
 * @summary Update a TrafficController
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/TrafficControllerPatch.json
 */
async function patchTrafficController() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "tc1";
  const properties = { tags: { key1: "value1" } };
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.trafficControllerInterface.update(
    resourceGroupName,
    trafficControllerName,
    properties
  );
  console.log(result);
}
