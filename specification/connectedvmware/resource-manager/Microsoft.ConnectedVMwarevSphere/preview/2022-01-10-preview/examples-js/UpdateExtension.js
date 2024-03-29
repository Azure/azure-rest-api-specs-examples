const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update the extension.
 *
 * @summary The operation to update the extension.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/UpdateExtension.json
 */
async function createOrUpdateAMachineExtensionPatch() {
  const subscriptionId = process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "myResourceGroup";
  const name = "myMachine";
  const extensionName = "CustomScriptExtension";
  const extensionParameters = {
    type: "CustomScriptExtension",
    publisher: "Microsoft.Compute",
    settings: {
      commandToExecute: 'powershell.exe -c "Get-Process | Where-Object { $_.CPU -lt 100 }"',
    },
    typeHandlerVersion: "1.10",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.machineExtensions.beginUpdateAndWait(
    resourceGroupName,
    name,
    extensionName,
    extensionParameters
  );
  console.log(result);
}
