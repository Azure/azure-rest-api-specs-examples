const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Release ownership of an existing virtual machine This operation can take a while to complete.
 *
 * @summary Release ownership of an existing virtual machine This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_UnClaim.json
 */
async function virtualMachinesUnClaim() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginUnClaimAndWait(resourceGroupName, labName, name);
  console.log(result);
}

virtualMachinesUnClaim().catch(console.error);
