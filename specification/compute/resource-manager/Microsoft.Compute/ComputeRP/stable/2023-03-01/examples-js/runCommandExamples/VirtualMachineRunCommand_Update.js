const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update the run command.
 *
 * @summary The operation to update the run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/runCommandExamples/VirtualMachineRunCommand_Update.json
 */
async function updateARunCommand() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
  const runCommandName = "myRunCommand";
  const runCommand = {
    asyncExecution: false,
    errorBlobManagedIdentity: {
      objectId: "4231e4d2-33e4-4e23-96b2-17888afa6072",
    },
    errorBlobUri: "https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt",
    outputBlobUri:
      "https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt?sp=racw&st=2022-10-07T19:40:21Z&se=2022-10-08T03:40:21Z&spr=https&sv=2021-06-08&sr=b&sig=Yh7B%2Fy83olbYBdfsfbUREvd7ol8Dq5EVP3lAO4Kj4xDcN8%3D",
    parameters: [
      { name: "param1", value: "value1" },
      { name: "param2", value: "value2" },
    ],
    runAsPassword: "<runAsPassword>",
    runAsUser: "user1",
    source: {
      script: "Write-Host Hello World! ; Remove-Item C:\test\testFile.txt",
    },
    timeoutInSeconds: 3600,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineRunCommands.beginUpdateAndWait(
    resourceGroupName,
    vmName,
    runCommandName,
    runCommand
  );
  console.log(result);
}
