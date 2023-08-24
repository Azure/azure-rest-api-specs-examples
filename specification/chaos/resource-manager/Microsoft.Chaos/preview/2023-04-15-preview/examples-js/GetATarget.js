const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Target resource that extends a tracked regional resource.
 *
 * @summary Get a Target resource that extends a tracked regional resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetATarget.json
 */
async function getATargetThatExtendsAVirtualMachineResource() {
  const subscriptionId =
    process.env["CHAOS_SUBSCRIPTION_ID"] || "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = process.env["CHAOS_RESOURCE_GROUP"] || "exampleRG";
  const parentProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResourceName = "exampleVM";
  const targetName = "Microsoft-Agent";
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.targets.get(
    resourceGroupName,
    parentProviderNamespace,
    parentResourceType,
    parentResourceName,
    targetName
  );
  console.log(result);
}
