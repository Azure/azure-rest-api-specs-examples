const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a capacity pool
 *
 * @summary Create or Update a capacity pool
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2024-07-01-preview/examples/Pools_CreateOrUpdate_CustomThroughput.json
 */
async function poolsCreateOrUpdateCustomThroughput() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const poolName = "customPool1";
  const body = {
    customThroughputMibps: 128,
    location: "eastus",
    qosType: "Manual",
    serviceLevel: "Flexible",
    size: 4398046511104,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.pools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    poolName,
    body,
  );
  console.log(result);
}
