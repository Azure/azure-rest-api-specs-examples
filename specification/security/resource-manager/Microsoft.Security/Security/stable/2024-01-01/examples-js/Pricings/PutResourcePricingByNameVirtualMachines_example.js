const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a provided Microsoft Defender for Cloud pricing configuration in the scope. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines' and only for plan='VirtualMachines' and subPlan='P1').
 *
 * @summary updates a provided Microsoft Defender for Cloud pricing configuration in the scope. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines' and only for plan='VirtualMachines' and subPlan='P1').
 * x-ms-original-file: 2024-01-01/Pricings/PutResourcePricingByNameVirtualMachines_example.json
 */
async function updatePricingOnResourceExampleForVirtualMachinesPlan() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.update(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1",
    "virtualMachines",
    { pricingTier: "Standard", subPlan: "P1" },
  );
  console.log(result);
}
