const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new cluster manager or update properties of the cluster manager if it exists.
 *
 * @summary Create a new cluster manager or update properties of the cluster manager if it exists.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/ClusterManagers_Create.json
 */
async function createOrUpdateClusterManager() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterManagerName = "clusterManagerName";
  const clusterManagerParameters = {
    analyticsWorkspaceId:
      "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName",
    fabricControllerId:
      "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName",
    location: "location",
    managedResourceGroupConfiguration: {
      name: "my-managed-rg",
      location: "East US",
    },
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.clusterManagers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterManagerName,
    clusterManagerParameters
  );
  console.log(result);
}
