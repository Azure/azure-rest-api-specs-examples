const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an HCI cluster.
 *
 * @summary Update an HCI cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/UpdateCluster.json
 */
async function updateCluster() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const cluster = {
    type: "SystemAssigned",
    cloudManagementEndpoint: "https://98294836-31be-4668-aeae-698667faf99b.waconazure.com",
    desiredProperties: {
      diagnosticLevel: "Basic",
      windowsServerSubscription: "Enabled",
    },
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.clusters.update(resourceGroupName, clusterName, cluster);
  console.log(result);
}
