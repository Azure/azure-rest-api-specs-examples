const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the IP Prefix resource.
 *
 * @summary API to update certain properties of the IP Prefix resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpPrefixes_Update_MaximumSet_Gen.json
 */
async function ipPrefixesUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const ipPrefixName = "example-ipPrefix";
  const body = {
    annotation: "annotation",
    ipPrefixRules: [
      {
        action: "Permit",
        condition: "GreaterThanOrEqualTo",
        networkPrefix: "10.10.10.10/30",
        sequenceNumber: 4155123341,
        subnetMaskLength: "10",
      },
    ],
    tags: { keyID: "KeyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipPrefixes.beginUpdateAndWait(resourceGroupName, ipPrefixName, body);
  console.log(result);
}
