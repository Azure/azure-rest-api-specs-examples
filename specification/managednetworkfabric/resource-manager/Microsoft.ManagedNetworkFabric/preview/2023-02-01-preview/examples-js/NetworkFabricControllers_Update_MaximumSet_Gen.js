const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates are currently not supported for the Network Fabric Controller resource.
 *
 * @summary Updates are currently not supported for the Network Fabric Controller resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_Update_MaximumSet_Gen.json
 */
async function networkFabricControllersUpdateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkFabricControllerName = "networkFabricControllerName";
  const body = {
    workloadExpressRouteConnections: [
      {
        expressRouteAuthorizationKey: "xxxxxxx",
        expressRouteCircuitId:
          "/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkFabricControllers.beginUpdateAndWait(
    resourceGroupName,
    networkFabricControllerName,
    body
  );
  console.log(result);
}
