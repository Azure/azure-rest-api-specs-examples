const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the L2 Isolation Domain resource.
 *
 * @summary API to update certain properties of the L2 Isolation Domain resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_Update_MaximumSet_Gen.json
 */
async function l2IsolationDomainsUpdateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l2IsolationDomainName = "l2IsolationDomainName";
  const body = { mtu: 9000 };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.l2IsolationDomains.beginUpdateAndWait(
    resourceGroupName,
    l2IsolationDomainName,
    body
  );
  console.log(result);
}
