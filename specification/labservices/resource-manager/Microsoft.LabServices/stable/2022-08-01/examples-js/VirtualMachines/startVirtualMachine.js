const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Action to start a lab virtual machine.
 *
 * @summary Action to start a lab virtual machine.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/startVirtualMachine.json
 */
async function startVirtualMachine() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const virtualMachineName = "template";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginStartAndWait(
    resourceGroupName,
    labName,
    virtualMachineName
  );
  console.log(result);
}

startVirtualMachine().catch(console.error);
