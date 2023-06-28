const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clears IPv6 neighbor tables for this Isolation Domain.
 *
 * @summary Clears IPv6 neighbor tables for this Isolation Domain.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L3IsolationDomains_clearNeighborTable_MaximumSet_Gen.json
 */
async function l3IsolationDomainsClearNeighborTableMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l3IsolationDomainName = "example-l3domain";
  const body = { resourceIds: ["wnaxqikneofcni"] };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.l3IsolationDomains.beginClearNeighborTableAndWait(
    resourceGroupName,
    l3IsolationDomainName,
    body
  );
  console.log(result);
}
