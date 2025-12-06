const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or Update a capacity pool
 *
 * @summary create or Update a capacity pool
 * x-ms-original-file: 2025-09-01-preview/Pools_CreateOrUpdate_CustomThroughput.json
 */
async function poolsCreateOrUpdateCustomThroughput() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.pools.createOrUpdate("myRG", "account1", "customPool1", {
    location: "eastus",
    properties: {
      customThroughputMibps: 128,
      qosType: "Manual",
      serviceLevel: "Flexible",
      size: 4398046511104,
    },
  });
  console.log(result);
}
