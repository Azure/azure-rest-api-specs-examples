const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a provided Microsoft Defender for Cloud pricing configuration in the scope. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines' and only for plan='VirtualMachines' and subPlan='P1').
 *
 * @summary updates a provided Microsoft Defender for Cloud pricing configuration in the scope. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines' and only for plan='VirtualMachines' and subPlan='P1').
 * x-ms-original-file: 2024-01-01/Pricings/PutResourcePricingByNameContainers_example.json
 */
async function updatePricingOnResourceExampleForContainersPlan() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.update(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/demo-containers-rg/providers/Microsoft.ContainerService/managedClusters/demo-aks-cluster",
    "Containers",
    {
      pricingTier: "Standard",
      extensions: [
        { name: "ContainerRegistriesVulnerabilityAssessments", isEnabled: "True" },
        { name: "ContainerSensor", isEnabled: "True" },
        { name: "AgentlessDiscoveryForKubernetes", isEnabled: "True" },
        {
          name: "AgentlessVmScanning",
          additionalExtensionProperties: { ExclusionTags: "[]" },
          isEnabled: "True",
        },
        { name: "ContainerIntegrityContribution", isEnabled: "True" },
      ],
    },
  );
  console.log(result);
}
