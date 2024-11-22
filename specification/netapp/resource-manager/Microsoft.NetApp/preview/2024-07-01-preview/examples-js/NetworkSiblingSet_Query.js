const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details of the specified network sibling set.
 *
 * @summary Get details of the specified network sibling set.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2024-07-01-preview/examples/NetworkSiblingSet_Query.json
 */
async function networkSiblingSetQuery() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const location = "eastus";
  const networkSiblingSetId = "9760acf5-4638-11e7-9bdb-020073ca3333";
  const subnetId =
    "/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.netAppResource.queryNetworkSiblingSet(
    location,
    networkSiblingSetId,
    subnetId,
  );
  console.log(result);
}
