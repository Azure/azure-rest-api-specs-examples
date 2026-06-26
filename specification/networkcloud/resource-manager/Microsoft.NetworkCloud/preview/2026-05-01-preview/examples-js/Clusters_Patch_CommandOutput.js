const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 *
 * @summary patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 * x-ms-original-file: 2026-05-01-preview/Clusters_Patch_CommandOutput.json
 */
async function patchClusterCommandOutput() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.clusters.update("resourceGroupName", "clusterName", {
    clusterUpdateParameters: {
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1":
            {},
          "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2":
            {},
        },
      },
      commandOutputSettings: {
        associatedIdentity: {
          identityType: "UserAssignedIdentity",
          userAssignedIdentityResourceId:
            "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2",
        },
        containerUrl: "https://myaccount.blob.core.windows.net/mycontainer?restype=container",
        overrides: [
          {
            associatedIdentity: {
              identityType: "UserAssignedIdentity",
              userAssignedIdentityResourceId:
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2",
            },
            commandOutputType: "StorageRunReadCommands",
            containerUrl: "https://myaccount.blob.core.windows.net/myContainer2?restype=container",
          },
        ],
      },
      tags: { key1: "myvalue1", key2: "myvalue2" },
    },
  });
  console.log(result);
}
