const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a provided Microsoft Defender for Cloud pricing configuration in a specific resource. Valid only for resource scope (Supported resources are: 'VirtualMachines, VMSS and ARC MachinesS').
 *
 * @summary Deletes a provided Microsoft Defender for Cloud pricing configuration in a specific resource. Valid only for resource scope (Supported resources are: 'VirtualMachines, VMSS and ARC MachinesS').
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/DeleteResourcePricing_example.json
 */
async function deleteAPricingOnResource() {
  const scopeId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1";
  const pricingName = "VirtualMachines";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.delete(scopeId, pricingName);
  console.log(result);
}
