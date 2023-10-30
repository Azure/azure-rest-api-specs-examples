const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the virtual machine instances within the specified parent resource.
 *
 * @summary Lists all of the virtual machine instances within the specified parent resource.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListVirtualMachineInstances.json
 */
async function listVirtualMachineInstances() {
  const resourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential);
  const resArray = new Array();
  for await (let item of client.virtualMachineInstances.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
