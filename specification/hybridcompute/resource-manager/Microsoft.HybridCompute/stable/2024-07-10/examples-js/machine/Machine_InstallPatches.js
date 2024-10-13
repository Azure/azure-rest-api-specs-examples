const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to install patches on a hybrid machine identity in Azure.
 *
 * @summary The operation to install patches on a hybrid machine identity in Azure.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/machine/Machine_InstallPatches.json
 */
async function installPatchStateOfAMachine() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroupName";
  const name = "myMachineName";
  const installPatchesInput = {
    maximumDuration: "PT4H",
    rebootSetting: "IfRequired",
    windowsParameters: {
      classificationsToInclude: ["Critical", "Security"],
      maxPatchPublishDate: new Date("2021-08-19T02:36:43.0539904+00:00"),
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machines.beginInstallPatchesAndWait(
    resourceGroupName,
    name,
    installPatchesInput,
  );
  console.log(result);
}
