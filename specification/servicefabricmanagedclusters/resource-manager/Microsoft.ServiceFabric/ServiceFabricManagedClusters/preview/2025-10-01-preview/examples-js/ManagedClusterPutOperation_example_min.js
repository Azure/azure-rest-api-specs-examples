const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric managed cluster resource with the specified name.
 *
 * @summary create or update a Service Fabric managed cluster resource with the specified name.
 * x-ms-original-file: 2025-10-01-preview/ManagedClusterPutOperation_example_min.json
 */
async function putAClusterWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("resRg", "myCluster", {
    location: "eastus",
    properties: {
      adminPassword: "{vm-password}",
      adminUserName: "vmadmin",
      clusterUpgradeCadence: "Wave1",
      clusterUpgradeMode: "Automatic",
      dnsName: "myCluster",
      fabricSettings: [
        {
          name: "ManagedIdentityTokenService",
          parameters: [{ name: "IsEnabled", value: "true" }],
        },
      ],
    },
    sku: { name: "Basic" },
  });
  console.log(result);
}
