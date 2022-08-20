const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get all extensions of a non-Azure machine
 *
 * @summary The operation to get all extensions of a non-Azure machine
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/LISTExtension.json
 */
async function getAllMachineExtensions() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const name = "myMachine";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.machineExtensions.list(resourceGroupName, name)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllMachineExtensions().catch(console.error);
