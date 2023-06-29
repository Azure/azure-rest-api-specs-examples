const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements IP Prefix DELETE method.
 *
 * @summary Implements IP Prefix DELETE method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpPrefixes_Delete_MaximumSet_Gen.json
 */
async function ipPrefixesDeleteMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "xxxxxx";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "rgIpPrefixLists";
  const ipPrefixName = "example-ipPrefix";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipPrefixes.beginDeleteAndWait(resourceGroupName, ipPrefixName);
  console.log(result);
}
