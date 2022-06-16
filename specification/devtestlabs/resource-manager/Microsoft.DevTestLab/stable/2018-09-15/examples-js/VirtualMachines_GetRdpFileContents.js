const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a string that represents the contents of the RDP file for the virtual machine
 *
 * @summary Gets a string that represents the contents of the RDP file for the virtual machine
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_GetRdpFileContents.json
 */
async function virtualMachinesGetRdpFileContents() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.getRdpFileContents(resourceGroupName, labName, name);
  console.log(result);
}

virtualMachinesGetRdpFileContents().catch(console.error);
