const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Traffic Controller Frontend
 *
 * @summary Update a Traffic Controller Frontend
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPatch.json
 */
async function updateFrontend() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "TC1";
  const frontendName = "publicIp1";
  const properties = {
    properties: {
      ipAddressVersion: "IPv4",
      mode: "public",
      publicIPAddress: { id: "resourceUriAsString" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.frontendsInterface.update(
    resourceGroupName,
    trafficControllerName,
    frontendName,
    properties
  );
  console.log(result);
}
