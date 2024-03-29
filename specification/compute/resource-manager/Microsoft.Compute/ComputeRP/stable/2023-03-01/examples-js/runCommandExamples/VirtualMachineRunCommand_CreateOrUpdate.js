const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update the run command.
 *
 * @summary The operation to create or update the run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/runCommandExamples/VirtualMachineRunCommand_CreateOrUpdate.json
 */
async function createOrUpdateARunCommand() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
  const runCommandName = "myRunCommand";
  const runCommand = {
    asyncExecution: false,
    errorBlobUri:
      "https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt?sp=racw&st=2022-10-07T19:40:21Z&se=2022-10-08T03:40:21Z&spr=https&sv=2021-06-08&sr=b&sig=Yh7B%2Fy83olbYBdfsfbUREvd7ol8Dq5EVP3lAO4Kj4xDcN8%3D",
    location: "West US",
    outputBlobManagedIdentity: {
      clientId: "22d35efb-0c99-4041-8c5b-6d24db33a69a",
    },
    outputBlobUri:
      "https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt",
    parameters: [
      { name: "param1", value: "value1" },
      { name: "param2", value: "value2" },
    ],
    runAsPassword: "<runAsPassword>",
    runAsUser: "user1",
    source: {
      scriptUri:
        "https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1?sp=r&st=2022-10-07T19:52:54Z&se=2022-10-08T03:52:54Z&spr=https&sv=2021-06-08&sr=b&sig=zfYFYCgea1PqVERZuwJiewrte5gjTnKGtVJngcw5oc828%3D",
    },
    timeoutInSeconds: 3600,
    treatFailureAsDeploymentFailure: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineRunCommands.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmName,
    runCommandName,
    runCommand
  );
  console.log(result);
}
