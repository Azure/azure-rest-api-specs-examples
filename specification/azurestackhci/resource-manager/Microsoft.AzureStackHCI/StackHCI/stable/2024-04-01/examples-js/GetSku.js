const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get SKU resource details within a offer of HCI Cluster.
 *
 * @summary Get SKU resource details within a offer of HCI Cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetSku.json
 */
async function getSku() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const publisherName = "publisher1";
  const offerName = "offer1";
  const skuName = "sku1";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.skus.get(
    resourceGroupName,
    clusterName,
    publisherName,
    offerName,
    skuName,
  );
  console.log(result);
}
