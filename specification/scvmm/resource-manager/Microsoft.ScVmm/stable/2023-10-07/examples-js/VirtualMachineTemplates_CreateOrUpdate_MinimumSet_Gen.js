const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the ScVmm VM Template as an Azure VM Template resource.
 *
 * @summary Onboards the ScVmm VM Template as an Azure VM Template resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_CreateOrUpdate_MinimumSet_Gen.json
 */
async function virtualMachineTemplatesCreateOrUpdateMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const virtualMachineTemplateName = "P";
  const resource = {
    extendedLocation: {},
    location: "ayxsyduviotylbojh",
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualMachineTemplateName,
    resource,
  );
  console.log(result);
}
