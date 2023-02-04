const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the VirtualMachineTemplate resource.
 *
 * @summary Updates the VirtualMachineTemplate resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/UpdateVirtualMachineTemplate.json
 */
async function updateVirtualMachineTemplate() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const virtualMachineTemplateName = "HRVirtualMachineTemplate";
  const body = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.beginUpdateAndWait(
    resourceGroupName,
    virtualMachineTemplateName,
    body
  );
  console.log(result);
}
