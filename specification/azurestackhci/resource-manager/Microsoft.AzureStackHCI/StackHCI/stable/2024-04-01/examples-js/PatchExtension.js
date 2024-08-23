const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update Extension for HCI cluster.
 *
 * @summary Update Extension for HCI cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PatchExtension.json
 */
async function updateArcExtension() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const extensionName = "MicrosoftMonitoringAgent";
  const extension = {
    extensionParameters: {
      enableAutomaticUpgrade: false,
      protectedSettings: { workspaceKey: "xx" },
      settings: { workspaceId: "xx" },
      typeHandlerVersion: "1.10",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.extensions.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    arcSettingName,
    extensionName,
    extension,
  );
  console.log(result);
}
