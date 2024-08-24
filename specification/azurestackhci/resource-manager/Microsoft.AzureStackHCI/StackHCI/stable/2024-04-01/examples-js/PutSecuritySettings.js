const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a security setting
 *
 * @summary Create a security setting
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutSecuritySettings.json
 */
async function createSecuritySettings() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const securitySettingsName = "default";
  const resource = {
    securedCoreComplianceAssignment: "Audit",
    smbEncryptionForIntraClusterTrafficComplianceAssignment: "Audit",
    wdacComplianceAssignment: "ApplyAndAutoCorrect",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.securitySettings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    securitySettingsName,
    resource,
  );
  console.log(result);
}
