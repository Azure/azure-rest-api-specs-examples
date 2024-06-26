const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Target resource that extends a tracked regional resource.
 *
 * @summary Delete a Target resource that extends a tracked regional resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/DeleteTarget.json
 */
async function deleteATargetThatExtendsAVirtualMachineResource() {
  const subscriptionId =
    process.env["CHAOS_SUBSCRIPTION_ID"] || "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = process.env["CHAOS_RESOURCE_GROUP"] || "exampleRG";
  const parentProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResourceName = "exampleVM";
  const targetName = "Microsoft-Agent";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.targets.delete(
    resourceGroupName,
    parentProviderNamespace,
    parentResourceType,
    parentResourceName,
    targetName,
  );
  console.log(result);
}
