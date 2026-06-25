const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_FromWithNetworkInterfaceWithDnsSettings.json
 */
async function createAScaleSetWithNetworkInterfacesWithPublicIpAddressDnsSettings() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 3, name: "Standard_D1_v2" },
      location: "westus",
      overprovision: true,
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}",
          },
          osDisk: {
            caching: "ReadWrite",
            managedDisk: { storageAccountType: "Standard_LRS" },
            createOption: "FromImage",
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
              name: "{nicConfig1-name}",
              tags: { nicTag: "tag" },
              primary: true,
              enableIPForwarding: true,
              disableTcpStateTracking: true,
              enableAcceleratedNetworking: true,
              auxiliaryMode: "AcceleratedConnections",
              auxiliarySku: "A1",
              ipConfigurations: [
                {
                  name: "{vmss-name}",
                  subnet: {
                    id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                  },
                },
              ],
            },
            {
              name: "{nicConfig2-name}",
              primary: false,
              enableAcceleratedNetworking: false,
              enableIPForwarding: false,
              disableTcpStateTracking: false,
              ipConfigurations: [
                {
                  name: "{nicConfig2-name}",
                  primary: true,
                  subnet: {
                    id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-fpga-subnet-name2}",
                  },
                  privateIPAddressVersion: "IPv4",
                  publicIPAddressConfiguration: {
                    name: "publicip",
                    tags: { pipTag: "tag" },
                    idleTimeoutInMinutes: 10,
                    dnsSettings: {
                      domainNameLabel: "vmsstestlabel01",
                      domainNameLabelScope: "NoReuse",
                    },
                  },
                },
              ],
            },
          ],
        },
      },
      upgradePolicy: { mode: "Manual" },
    },
  );
  console.log(result);
}
