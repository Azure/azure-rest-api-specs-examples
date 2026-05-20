const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new HDInsight cluster with the specified parameters.
 *
 * @summary creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: 2025-01-15-preview/CreateHDInsightClusterWithADLSGen2Msi.json
 */
async function createClusterWithStorageAdlsGen2MSI() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.create("rg1", "cluster1", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi":
          {},
      },
    },
    properties: {
      clusterDefinition: {
        configurations: {
          gateway: {
            "restAuthCredential.isEnabled": true,
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
        },
        kind: "Hadoop",
      },
      clusterVersion: "5.1",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "Standard_E8_V3" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: { password: "**********", username: "sshuser" },
            },
            scriptActions: [],
            targetInstanceCount: 2,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
          {
            name: "workernode",
            hardwareProfile: { vmSize: "Standard_E8_V3" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: { password: "**********", username: "sshuser" },
            },
            scriptActions: [],
            targetInstanceCount: 3,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
          {
            name: "zookeepernode",
            hardwareProfile: { vmSize: "Standard_E8_V3" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: { password: "**********", username: "sshuser" },
            },
            scriptActions: [],
            targetInstanceCount: 3,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
        ],
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage.blob.core.windows.net",
            fileSystem: "default",
            isDefault: true,
            msiResourceId:
              "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi",
            resourceId:
              "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/mystorage",
          },
        ],
      },
      tier: "Standard",
    },
    tags: { key1: "val1" },
  });
  console.log(result);
}
