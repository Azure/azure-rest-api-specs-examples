const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deregisters the ScVmm VM Template from Azure.
 *
 * @summary Deregisters the ScVmm VM Template from Azure.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVirtualMachineTemplate.json
 */
async function deleteVirtualMachineTemplate() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const virtualMachineTemplateName = "HRVirtualMachineTemplate";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.beginDeleteAndWait(
    resourceGroupName,
    virtualMachineTemplateName
  );
  console.log(result);
}
