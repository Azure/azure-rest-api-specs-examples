const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an express route circuit.
 *
 * @summary creates or updates an express route circuit.
 * x-ms-original-file: 2025-09-01/ExpressRouteMultiCloudCircuitCreateWithPartnerAccountId.json
 */
async function createMultiCloudExpressRouteCircuitWithPartnerAccountId() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.createOrUpdate("rg1", "circuitName", {
    serviceProviderProperties: {
      serviceProviderName: "AWS",
      peeringLocation: "uswest2",
      bandwidthInMbps: 200,
    },
    partnerAccountId: "123456789",
    sku: { name: "MultiCloud_MeteredData", tier: "MultiCloud", family: "MeteredData" },
    location: "eastus2euap",
  });
  console.log(result);
}
