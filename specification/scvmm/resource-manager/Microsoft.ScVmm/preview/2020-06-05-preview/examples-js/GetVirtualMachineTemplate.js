const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements VirtualMachineTemplate GET method.
 *
 * @summary Implements VirtualMachineTemplate GET method.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetVirtualMachineTemplate.json
 */
async function getVirtualMachineTemplate() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const virtualMachineTemplateName = "HRVirtualMachineTemplate";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.get(
    resourceGroupName,
    virtualMachineTemplateName
  );
  console.log(result);
}

getVirtualMachineTemplate().catch(console.error);
