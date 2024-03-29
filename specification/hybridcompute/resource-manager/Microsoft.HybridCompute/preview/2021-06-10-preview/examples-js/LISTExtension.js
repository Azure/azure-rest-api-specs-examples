const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get all extensions of a non-Azure machine
 *
 * @summary The operation to get all extensions of a non-Azure machine
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/LISTExtension.json
 */
async function getAllMachineExtensions() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.machineExtensions.list(resourceGroupName, machineName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllMachineExtensions().catch(console.error);
