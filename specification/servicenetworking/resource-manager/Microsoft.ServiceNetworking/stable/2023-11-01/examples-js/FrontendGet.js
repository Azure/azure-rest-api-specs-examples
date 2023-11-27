const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Frontend
 *
 * @summary Get a Frontend
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/FrontendGet.json
 */
async function getFrontend() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "tc1";
  const frontendName = "fe1";
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.frontendsInterface.get(
    resourceGroupName,
    trafficControllerName,
    frontendName
  );
  console.log(result);
}
