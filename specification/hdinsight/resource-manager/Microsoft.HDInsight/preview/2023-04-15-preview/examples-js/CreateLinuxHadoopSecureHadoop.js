const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateLinuxHadoopSecureHadoop.json
 */
async function createSecureHadoopCluster() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
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
      clusterVersion: "3.5",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "Standard_D3_V2" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            scriptActions: [],
            targetInstanceCount: 2,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
          {
            name: "workernode",
            hardwareProfile: { vmSize: "Standard_D3_V2" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            scriptActions: [],
            targetInstanceCount: 4,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
          {
            name: "zookeepernode",
            hardwareProfile: { vmSize: "Small" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            scriptActions: [],
            targetInstanceCount: 3,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
        ],
      },
      osType: "Linux",
      securityProfile: {
        clusterUsersGroupDNs: ["hdiusers"],
        directoryType: "ActiveDirectory",
        domain: "DomainName",
        domainUserPassword: "**********",
        domainUsername: "DomainUsername",
        ldapsUrls: ["ldaps://10.10.0.4:636"],
        organizationalUnitDN: "OU=Hadoop,DC=hdinsight,DC=test",
      },
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage.blob.core.windows.net",
            container: "containername",
            enableSecureChannel: true,
            isDefault: true,
            key: "storage account key",
          },
        ],
      },
      tier: "Premium",
    },
    tags: { key1: "val1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}
