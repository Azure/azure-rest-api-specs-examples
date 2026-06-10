const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the extension.
 *
 * @summary the operation to get the extension.
 * x-ms-original-file: 2025-09-16-preview/extension/Extension_Get.json
 */
async function getMachineExtension() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machineExtensions.get(
    "myResourceGroup",
    "myMachine",
    "CustomScriptExtension",
  );
  console.log(result);
}
