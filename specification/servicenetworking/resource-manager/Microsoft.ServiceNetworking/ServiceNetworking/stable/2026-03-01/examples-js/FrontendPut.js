const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Frontend
 *
 * @summary create a Frontend
 * x-ms-original-file: 2026-03-01/FrontendPut.json
 */
async function putFrontend() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.frontendsInterface.createOrUpdate("rg1", "tc1", "fe1", {
    location: "NorthCentralUS",
    properties: { publicNetworkAccess: "Enabled" },
  });
  console.log(result);
}
