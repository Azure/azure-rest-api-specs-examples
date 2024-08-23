const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update ArcSettings for HCI cluster.
 *
 * @summary Update ArcSettings for HCI cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PatchArcSetting.json
 */
async function patchArcSetting() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const arcSetting = {
    connectivityProperties: {
      enabled: true,
      serviceConfigurations: [{ port: 6516, serviceName: "WAC" }],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.arcSettings.update(
    resourceGroupName,
    clusterName,
    arcSettingName,
    arcSetting,
  );
  console.log(result);
}
