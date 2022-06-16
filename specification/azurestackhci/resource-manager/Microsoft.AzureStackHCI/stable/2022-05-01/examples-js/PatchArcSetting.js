const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update ArcSettings for HCI cluster.
 *
 * @summary Update ArcSettings for HCI cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/PatchArcSetting.json
 */
async function patchArcSetting() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const arcSetting = {
    connectivityProperties: { enabled: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.arcSettings.update(
    resourceGroupName,
    clusterName,
    arcSettingName,
    arcSetting
  );
  console.log(result);
}

patchArcSetting().catch(console.error);
