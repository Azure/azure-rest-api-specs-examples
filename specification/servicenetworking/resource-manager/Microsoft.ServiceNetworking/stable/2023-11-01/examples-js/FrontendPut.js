const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Frontend
 *
 * @summary Create a Frontend
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/FrontendPut.json
 */
async function putFrontend() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "tc1";
  const frontendName = "fe1";
  const resource = { location: "NorthCentralUS", properties: {} };
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.frontendsInterface.beginCreateOrUpdateAndWait(
    resourceGroupName,
    trafficControllerName,
    frontendName,
    resource
  );
  console.log(result);
}
