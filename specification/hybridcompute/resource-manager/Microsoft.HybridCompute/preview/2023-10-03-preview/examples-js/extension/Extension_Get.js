const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the extension.
 *
 * @summary The operation to get the extension.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/extension/Extension_Get.json
 */
async function getMachineExtension() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const machineName = "myMachine";
  const extensionName = "CustomScriptExtension";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machineExtensions.get(resourceGroupName, machineName, extensionName);
  console.log(result);
}
