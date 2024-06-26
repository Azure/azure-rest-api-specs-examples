const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Run and retrieve output from read only commands on the provided storage appliance.
 *
 * @summary Run and retrieve output from read only commands on the provided storage appliance.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/StorageAppliances_RunReadCommands.json
 */
async function runAndRetrieveOutputFromReadOnlyCommandsOnStorageAppliance() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const storageApplianceName = "storageApplianceName";
  const storageApplianceRunReadCommandsParameters = {
    limitTimeSeconds: 60,
    commands: [{ command: "AlertList" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.storageAppliances.beginRunReadCommandsAndWait(
    resourceGroupName,
    storageApplianceName,
    storageApplianceRunReadCommandsParameters
  );
  console.log(result);
}
