const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSpotPlusPriorityFlex.json
 */
async function createAScaleSetWithSpotPlusPriorityFlexible() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { capacity: 10, name: "Mix" },
      location: "westus",
      orchestrationMode: "Flexible",
      platformFaultDomainCount: 1,
      singlePlacementGroup: false,
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            publisher: "Canonical",
            offer: "0001-com-ubuntu-server-focal",
            sku: "20_04-lts-gen2",
            version: "latest",
          },
          osDisk: {
            createOption: "FromImage",
            caching: "ReadWrite",
            managedDisk: { storageAccountType: "Standard_LRS" },
          },
        },
        osProfile: { computerNamePrefix: "{vmss-name}", adminUsername: "{your-username}" },
        networkProfile: {
          networkInterfaceConfigurations: [
            {
              name: "{vmss-name}",
              primary: true,
              enableIPForwarding: true,
              enableAcceleratedNetworking: false,
              ipConfigurations: [
                {
                  name: "{vmss-name}",
                  subnet: {
                    id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                  },
                  primary: true,
                  applicationGatewayBackendAddressPools: [],
                  loadBalancerBackendAddressPools: [],
                  publicIPAddressConfiguration: { name: "{vmss-name}", idleTimeoutInMinutes: 15 },
                },
              ],
            },
          ],
          networkApiVersion: "2020-11-01",
        },
        priority: "SpotPlus",
        evictionPolicy: "Deallocate",
        billingProfile: { maxPrice: -1 },
      },
      priorityMixPolicy: { baseRegularPriorityCount: 4, regularPriorityPercentageAboveBase: 50 },
      skuProfile: {
        vmSizes: [
          { name: "Standard_D8s_v5" },
          { name: "Standard_E16s_v5" },
          { name: "Standard_D2s_v5" },
        ],
        allocationStrategy: "CapacityOptimized",
      },
    },
  );
  console.log(result);
}
