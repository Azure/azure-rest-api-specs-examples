const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deregisters the ScVmm VM Template from Azure.
 *
 * @summary Deregisters the ScVmm VM Template from Azure.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Delete_MinimumSet_Gen.json
 */
async function virtualMachineTemplatesDeleteMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const virtualMachineTemplateName = "5";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.beginDeleteAndWait(
    resourceGroupName,
    virtualMachineTemplateName,
  );
  console.log(result);
}
