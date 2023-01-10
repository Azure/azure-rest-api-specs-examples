const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update the extension.
 *
 * @summary The operation to create or update the extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExamples/VirtualMachineExtensions_CreateOrUpdate_MinimumSet_Gen.json
 */
async function virtualMachineExtensionsCreateOrUpdateMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmName = "aaaa";
  const vmExtensionName = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const extensionParameters = { location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmName,
    vmExtensionName,
    extensionParameters
  );
  console.log(result);
}
