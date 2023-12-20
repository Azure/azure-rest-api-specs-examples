const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update an extension.
 *
 * @summary The operation to update an extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Update_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionUpdateMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const vmssExtensionName = "aaaa";
  const extensionParameters = {
    typePropertiesType: "{extension-Type}",
    autoUpgradeMinorVersion: true,
    enableAutomaticUpgrade: true,
    forceUpdateTag: "aaaaaaaaa",
    protectedSettings: {},
    provisionAfterExtensions: ["aa"],
    publisher: "{extension-Publisher}",
    settings: {},
    suppressFailures: true,
    typeHandlerVersion: "{handler-version}",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.beginUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName,
    extensionParameters,
  );
  console.log(result);
}
