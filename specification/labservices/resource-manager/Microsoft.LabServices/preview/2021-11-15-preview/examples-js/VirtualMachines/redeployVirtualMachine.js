const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Action to redeploy a lab virtual machine to a different compute node. For troubleshooting connectivity.
 *
 * @summary Action to redeploy a lab virtual machine to a different compute node. For troubleshooting connectivity.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/VirtualMachines/redeployVirtualMachine.json
 */
async function redeployVirtualMachine() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const virtualMachineName = "template";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginRedeployAndWait(
    resourceGroupName,
    labName,
    virtualMachineName
  );
  console.log(result);
}

redeployVirtualMachine().catch(console.error);
