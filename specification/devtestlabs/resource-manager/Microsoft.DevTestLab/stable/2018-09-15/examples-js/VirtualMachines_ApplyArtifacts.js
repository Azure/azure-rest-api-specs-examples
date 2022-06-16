const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Apply artifacts to virtual machine. This operation can take a while to complete.
 *
 * @summary Apply artifacts to virtual machine. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_ApplyArtifacts.json
 */
async function virtualMachinesApplyArtifacts() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const applyArtifactsRequest = {
    artifacts: [
      {
        artifactId:
          "/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/public repo/artifacts/windows-restart",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginApplyArtifactsAndWait(
    resourceGroupName,
    labName,
    name,
    applyArtifactsRequest
  );
  console.log(result);
}

virtualMachinesApplyArtifacts().catch(console.error);
