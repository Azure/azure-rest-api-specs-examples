const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clears ARP tables for this Isolation Domain.
 *
 * @summary Clears ARP tables for this Isolation Domain.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_clearArpTable_MaximumSet_Gen.json
 */
async function l2IsolationDomainsClearArpTableMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l2IsolationDomainName = "example-l2domain";
  const body = {
    resourceIds: [
      "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/example-l2domain",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.l2IsolationDomains.beginClearArpTableAndWait(
    resourceGroupName,
    l2IsolationDomainName,
    body
  );
  console.log(result);
}
