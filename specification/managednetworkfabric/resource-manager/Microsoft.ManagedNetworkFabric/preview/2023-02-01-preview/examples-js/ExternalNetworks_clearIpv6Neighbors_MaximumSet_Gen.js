const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to clearIpv6Neighbors for externalNetwork.
 *
 * @summary clearIpv6Neighbors for externalNetwork.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_clearIpv6Neighbors_MaximumSet_Gen.json
 */
async function externalNetworksClearIpv6NeighborsMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l3IsolationDomainName = "example-l3domain";
  const externalNetworkName = "example-externalnetwork";
  const body = {
    resourceIds: [
      "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.externalNetworks.beginClearIpv6NeighborsAndWait(
    resourceGroupName,
    l3IsolationDomainName,
    externalNetworkName,
    body
  );
  console.log(result);
}
