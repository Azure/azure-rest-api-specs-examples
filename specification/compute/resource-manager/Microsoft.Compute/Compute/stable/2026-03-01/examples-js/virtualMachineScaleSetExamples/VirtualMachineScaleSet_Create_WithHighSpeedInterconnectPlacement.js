const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithHighSpeedInterconnectPlacement.json
 */
async function createAVirtualMachineScaleSetWithHighSpeedInterconnectPlacement() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 2, name: "Standard_D1_v2" },
      location: "westus",
      orchestrationMode: "Flexible",
      highSpeedInterconnectPlacement: "None",
      platformFaultDomainCount: 1,
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
        osProfile: {
          computerNamePrefix: "{vmss-name}",
          adminUsername: "{your-username}",
          adminPassword: "{your-password}",
        },
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
      },
    },
  );
  console.log(result);
}
