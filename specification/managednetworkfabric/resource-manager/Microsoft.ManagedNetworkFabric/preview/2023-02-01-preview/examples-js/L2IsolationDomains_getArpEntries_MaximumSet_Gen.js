const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clears IPv6 neighbors for this Isolation Domain.
 *
 * @summary Clears IPv6 neighbors for this Isolation Domain.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_getArpEntries_MaximumSet_Gen.json
 */
async function l2IsolationDomainsGetArpEntriesMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l2IsolationDomainName = "l2IsolationDomainName";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.l2IsolationDomains.beginGetArpEntriesAndWait(
    resourceGroupName,
    l2IsolationDomainName
  );
  console.log(result);
}
