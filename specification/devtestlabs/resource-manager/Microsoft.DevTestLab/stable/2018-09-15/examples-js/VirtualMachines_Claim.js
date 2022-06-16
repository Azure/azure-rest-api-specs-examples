const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Take ownership of an existing virtual machine This operation can take a while to complete.
 *
 * @summary Take ownership of an existing virtual machine This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Claim.json
 */
async function virtualMachinesClaim() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginClaimAndWait(resourceGroupName, labName, name);
  console.log(result);
}

virtualMachinesClaim().catch(console.error);
