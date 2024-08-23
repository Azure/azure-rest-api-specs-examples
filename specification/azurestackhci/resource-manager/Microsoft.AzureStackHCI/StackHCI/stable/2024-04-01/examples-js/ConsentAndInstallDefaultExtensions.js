const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add consent time for default extensions and initiate extensions installation
 *
 * @summary Add consent time for default extensions and initiate extensions installation
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ConsentAndInstallDefaultExtensions.json
 */
async function consentAndInstallDefaultExtensions() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.arcSettings.consentAndInstallDefaultExtensions(
    resourceGroupName,
    clusterName,
    arcSettingName,
  );
  console.log(result);
}
