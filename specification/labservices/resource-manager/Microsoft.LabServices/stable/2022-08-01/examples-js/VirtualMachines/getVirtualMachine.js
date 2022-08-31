const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the properties for a lab virtual machine.
 *
 * @summary Returns the properties for a lab virtual machine.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/getVirtualMachine.json
 */
async function getVirtualMachine() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const virtualMachineName = "template";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.virtualMachines.get(resourceGroupName, labName, virtualMachineName);
  console.log(result);
}

getVirtualMachine().catch(console.error);
