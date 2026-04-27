const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DevicePool
 *
 * @summary create a DevicePool
 * x-ms-original-file: 2026-04-01-preview/DevicePools_CreateOrUpdate_MaximumSet_Gen.json
 */
async function devicePoolsCreateOrUpdateMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.devicePools.createOrUpdate("ArcInstance-rg", "devicePool1", {
    properties: {},
    location: "eastus",
  });
  console.log(result);
}
