const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Network Fabric Controller.
 *
 * @summary Creates a Network Fabric Controller.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_Create_MaximumSet_Gen.json
 */
async function networkFabricControllersCreateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkFabricControllerName = "NetworkControllerName";
  const body = {
    annotation: "lab 1",
    infrastructureExpressRouteConnections: [
      {
        expressRouteAuthorizationKey: "xxxxxxx",
        expressRouteCircuitId:
          "/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName",
      },
    ],
    ipv4AddressSpace: "172.253.0.0/19",
    location: "eastus",
    managedResourceGroupConfiguration: {
      name: "managedResourceGroupName",
      location: "eastus",
    },
    workloadExpressRouteConnections: [
      {
        expressRouteAuthorizationKey: "xxxxx",
        expressRouteCircuitId:
          "/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkFabricControllers.beginCreateAndWait(
    resourceGroupName,
    networkFabricControllerName,
    body
  );
  console.log(result);
}
