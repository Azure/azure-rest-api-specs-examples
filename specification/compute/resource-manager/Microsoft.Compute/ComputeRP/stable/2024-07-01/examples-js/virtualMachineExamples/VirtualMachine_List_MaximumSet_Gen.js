const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all of the virtual machines in the specified resource group. Use the nextLink property in the response to get the next page of virtual machines.
 *
 * @summary Lists all of the virtual machines in the specified resource group. Use the nextLink property in the response to get the next page of virtual machines.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/virtualMachineExamples/VirtualMachine_List_MaximumSet_Gen.json
 */
async function virtualMachineListMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const filter = "aaaaaaaaaaaaaaaaaaaaaaa";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachines.list(resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
