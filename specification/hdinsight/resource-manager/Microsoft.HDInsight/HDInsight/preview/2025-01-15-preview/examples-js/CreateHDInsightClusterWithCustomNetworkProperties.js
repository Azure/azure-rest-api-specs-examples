const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new HDInsight cluster with the specified parameters.
 *
 * @summary creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: 2025-01-15-preview/CreateHDInsightClusterWithCustomNetworkProperties.json
 */
async function createClusterWithNetworkProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.create("rg1", "cluster1", {
    properties: {
      clusterDefinition: {
        configurations: {
          gateway: {
            "restAuthCredential.isEnabled": true,
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
        },
        kind: "hadoop",
      },
      clusterVersion: "3.6",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "standard_d3" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
          {
            name: "workernode",
            hardwareProfile: { vmSize: "standard_d3" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
        ],
      },
      networkProperties: {
        privateLink: "Enabled",
        publicIpTag: { ipTagType: "FirstPartyUsage", tag: "/<TagName>" },
        resourceProviderConnection: "Outbound",
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage",
            container: "containername",
            enableSecureChannel: true,
            isDefault: true,
            key: "storage account key",
          },
        ],
      },
    },
  });
  console.log(result);
}
