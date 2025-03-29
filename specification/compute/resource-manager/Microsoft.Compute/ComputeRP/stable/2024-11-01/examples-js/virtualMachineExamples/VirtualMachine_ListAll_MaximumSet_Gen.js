const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all of the virtual machines in the specified subscription. Use the nextLink property in the response to get the next page of virtual machines.
 *
 * @summary Lists all of the virtual machines in the specified subscription. Use the nextLink property in the response to get the next page of virtual machines.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineExamples/VirtualMachine_ListAll_MaximumSet_Gen.json
 */
async function virtualMachineListAllMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const statusOnly = "aaaaaa";
  const filter = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const options = { statusOnly, filter };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualMachines.listAll(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
