const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric managed cluster resource with the specified name.
 *
 * @summary Create or update a Service Fabric managed cluster resource with the specified name.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ManagedClusterPutOperation_example_min.json
 */
async function putAClusterWithMinimumParameters() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const parameters = {
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
    location: "eastus",
    sku: { name: "Basic" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.managedClusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    parameters,
  );
  console.log(result);
}
