const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Converts SinglePlacementGroup property to false for a existing virtual machine scale set.
 *
 * @summary Converts SinglePlacementGroup property to false for a existing virtual machine scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetConvertToSinglePlacementGroupMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaa";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.convertToSinglePlacementGroup(
    resourceGroupName,
    vmScaleSetName,
    parameters,
  );
  console.log(result);
}
