const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the virtual machines in the specified subscription. Use the nextLink property in the response to get the next page of virtual machines.
 *
 * @summary Lists all of the virtual machines in the specified subscription. Use the nextLink property in the response to get the next page of virtual machines.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_ListAll_MaximumSet_Gen.json
 */
async function virtualMachinesListAllMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const statusOnly = "aaaaaa";
  const filter = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const options = { statusOnly, filter };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachines.listAll(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachinesListAllMaximumSetGen().catch(console.error);
