const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithHighSpeedInterconnectPlacement.json
 */
async function createAVirtualMachineScaleSetWithHighSpeedInterconnectPlacement() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const parameters = {
    highSpeedInterconnectPlacement: "None",
    location: "westus",
    orchestrationMode: "Flexible",
    platformFaultDomainCount: 1,
    sku: { name: "Standard_D1_v2", capacity: 2, tier: "Standard" },
    virtualMachineProfile: {
      networkProfile: {
        networkApiVersion: "2020-11-01",
        networkInterfaceConfigurations: [
          {
            name: "{vmss-name}",
            enableAcceleratedNetworking: false,
            enableIPForwarding: true,
            ipConfigurations: [
              {
                name: "{vmss-name}",
                applicationGatewayBackendAddressPools: [],
                loadBalancerBackendAddressPools: [],
                primary: true,
                publicIPAddressConfiguration: {
                  name: "{vmss-name}",
                  idleTimeoutInMinutes: 15,
                },
                subnet: {
                  id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                },
              },
            ],
            primary: true,
          },
        ],
      },
      osProfile: {
        adminPassword: "{your-password}",
        adminUsername: "{your-username}",
        computerNamePrefix: "{vmss-name}",
      },
      storageProfile: {
        imageReference: {
          offer: "0001-com-ubuntu-server-focal",
          publisher: "Canonical",
          sku: "20_04-lts-gen2",
          version: "latest",
        },
        osDisk: {
          caching: "ReadWrite",
          createOption: "FromImage",
          managedDisk: { storageAccountType: "Standard_LRS" },
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    parameters,
  );
  console.log(result);
}
