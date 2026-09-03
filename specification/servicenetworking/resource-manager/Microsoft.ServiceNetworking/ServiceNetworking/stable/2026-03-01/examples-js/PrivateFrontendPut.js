const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Frontend
 *
 * @summary create a Frontend
 * x-ms-original-file: 2026-03-01/PrivateFrontendPut.json
 */
async function putPrivateFrontend() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.frontendsInterface.createOrUpdate("rg1", "tc1", "pfe1", {
    location: "NorthCentralUS",
    properties: {
      publicNetworkAccess: "Disabled",
      association: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/associations/as1",
      },
    },
  });
  console.log(result);
}
