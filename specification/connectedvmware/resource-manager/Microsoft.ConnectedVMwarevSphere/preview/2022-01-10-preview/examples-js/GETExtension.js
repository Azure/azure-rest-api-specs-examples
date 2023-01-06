const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the extension.
 *
 * @summary The operation to get the extension.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GETExtension.json
 */
async function getMachineExtension() {
  const subscriptionId = process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "myResourceGroup";
  const name = "myMachine";
  const extensionName = "CustomScriptExtension";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.machineExtensions.get(resourceGroupName, name, extensionName);
  console.log(result);
}
