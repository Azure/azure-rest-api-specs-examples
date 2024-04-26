const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to Upgrade Machine Extensions.
 *
 * @summary The operation to Upgrade Machine Extensions.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/extension/Extensions_Upgrade.json
 */
async function upgradeMachineExtensions() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const machineName = "myMachine";
  const extensionUpgradeParameters = {
    extensionTargets: {
      microsoftAzureMonitoring: { targetVersion: "2.0" },
      microsoftComputeCustomScriptExtension: { targetVersion: "1.10" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.beginUpgradeExtensionsAndWait(
    resourceGroupName,
    machineName,
    extensionUpgradeParameters,
  );
  console.log(result);
}
