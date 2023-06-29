const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update BfdForStaticRoutes for internalNetwork.
 *
 * @summary Update BfdForStaticRoutes for internalNetwork.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_updateBfdForStaticRouteAdministrativeState_MaximumSet_Gen.json
 */
async function internalNetworksUpdateBfdForStaticRouteAdministrativeStateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l3IsolationDomainName = "example-l3domain";
  const internalNetworkName = "example-internalnetwork";
  const body = {
    resourceIds: [
      "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/example-l3domain/internalNetworks/example-internalnetwork",
    ],
    state: "Enable",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result =
    await client.internalNetworks.beginUpdateBfdForStaticRouteAdministrativeStateAndWait(
      resourceGroupName,
      l3IsolationDomainName,
      internalNetworkName,
      body
    );
  console.log(result);
}
