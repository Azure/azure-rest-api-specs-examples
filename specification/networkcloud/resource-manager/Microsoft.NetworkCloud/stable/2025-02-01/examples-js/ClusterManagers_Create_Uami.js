const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a new cluster manager or update properties of the cluster manager if it exists.
 *
 * @summary Create a new cluster manager or update properties of the cluster manager if it exists.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/ClusterManagers_Create_Uami.json
 */
async function createOrUpdateClusterManagerWithUserAssignedIdentity() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterManagerName = "clusterManagerName";
  const clusterManagerParameters = {
    analyticsWorkspaceId:
      "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName",
    fabricControllerId:
      "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/123e4567E89b12d3A456426655440000/resourceGroups/resourceGroupName/providers/MicrosoftManagedIdentity/userAssignedIdentities/userIdentity1":
          {},
      },
    },
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
    clusterManagerParameters,
  );
  console.log(result);
}
