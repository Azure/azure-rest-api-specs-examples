const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of all virtual machines for a lab.
 *
 * @summary Returns a list of all virtual machines for a lab.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/listVirtualMachine.json
 */
async function listVirtualMachine() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachines.listByLab(resourceGroupName, labName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualMachine().catch(console.error);
