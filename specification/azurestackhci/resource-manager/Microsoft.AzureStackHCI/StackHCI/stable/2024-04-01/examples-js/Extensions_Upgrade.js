const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upgrade a particular Arc Extension of HCI Cluster.
 *
 * @summary Upgrade a particular Arc Extension of HCI Cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/Extensions_Upgrade.json
 */
async function upgradeMachineExtensions() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const extensionName = "MicrosoftMonitoringAgent";
  const extensionUpgradeParameters = {
    targetVersion: "1.0.18062.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.extensions.beginUpgradeAndWait(
    resourceGroupName,
    clusterName,
    arcSettingName,
    extensionName,
    extensionUpgradeParameters,
  );
  console.log(result);
}
