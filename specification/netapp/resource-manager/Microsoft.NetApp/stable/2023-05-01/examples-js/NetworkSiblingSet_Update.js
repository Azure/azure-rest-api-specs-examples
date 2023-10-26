const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the network features of the specified network sibling set.
 *
 * @summary Update the network features of the specified network sibling set.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/NetworkSiblingSet_Update.json
 */
async function networkFeaturesUpdate() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const location = "eastus";
  const networkSiblingSetId = "9760acf5-4638-11e7-9bdb-020073ca3333";
  const subnetId =
    "/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet";
  const networkSiblingSetStateId = "12345_44420.8001578125";
  const networkFeatures = "Standard";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.netAppResource.beginUpdateNetworkSiblingSetAndWait(
    location,
    networkSiblingSetId,
    subnetId,
    networkSiblingSetStateId,
    networkFeatures
  );
  console.log(result);
}
