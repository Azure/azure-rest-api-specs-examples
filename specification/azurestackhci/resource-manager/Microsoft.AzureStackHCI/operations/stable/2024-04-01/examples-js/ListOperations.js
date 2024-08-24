const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all available Microsoft.AzureStackHCI provider operations
 *
 * @summary List all available Microsoft.AzureStackHCI provider operations
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/operations/stable/2024-04-01/examples/ListOperations.json
 */
async function createCluster() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.operations.list();
  console.log(result);
}
