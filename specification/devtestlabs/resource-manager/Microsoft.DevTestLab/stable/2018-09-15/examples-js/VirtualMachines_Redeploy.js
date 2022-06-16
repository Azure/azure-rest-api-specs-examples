const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Redeploy a virtual machine This operation can take a while to complete.
 *
 * @summary Redeploy a virtual machine This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Redeploy.json
 */
async function virtualMachinesRedeploy() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginRedeployAndWait(
    resourceGroupName,
    labName,
    name
  );
  console.log(result);
}

virtualMachinesRedeploy().catch(console.error);
