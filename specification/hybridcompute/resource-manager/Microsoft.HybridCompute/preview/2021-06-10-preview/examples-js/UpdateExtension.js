const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update the extension.
 *
 * @summary The operation to create or update the extension.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/UpdateExtension.json
 */
async function createOrUpdateAMachineExtension() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const extensionName = "CustomScriptExtension";
  const extensionParameters = {
    properties: {
      type: "CustomScriptExtension",
      publisher: "Microsoft.Compute",
      settings: {
        commandToExecute: 'powershell.exe -c "Get-Process | Where-Object { $_.CPU -lt 100 }"',
      },
      typeHandlerVersion: "1.10",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machineExtensions.beginUpdateAndWait(
    resourceGroupName,
    machineName,
    extensionName,
    extensionParameters
  );
  console.log(result);
}

createOrUpdateAMachineExtension().catch(console.error);
