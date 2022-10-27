const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithVMsInDifferentZones.json
 */
async function createAScaleSetWithVirtualMachinesInDifferentZones() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const options = {
    body: {
      location: "centralus",
      properties: {
        overprovision: true,
        upgradePolicy: { mode: "Automatic" },
        virtualMachineProfile: {
          networkProfile: {
            networkInterfaceConfigurations: [
              {
                name: "{vmss-name}",
                properties: {
                  enableIPForwarding: true,
                  ipConfigurations: [
                    {
                      name: "{vmss-name}",
                      properties: {
                        subnet: {
                          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                        },
                      },
                    },
                  ],
                  primary: true,
                },
              },
            ],
          },
          osProfile: {
            adminPassword: "{your-password}",
            adminUsername: "{your-username}",
            computerNamePrefix: "{vmss-name}",
          },
          storageProfile: {
            dataDisks: [
              { createOption: "Empty", diskSizeGB: 1023, lun: 0 },
              { createOption: "Empty", diskSizeGB: 1023, lun: 1 },
            ],
            imageReference: {
              offer: "WindowsServer",
              publisher: "MicrosoftWindowsServer",
              sku: "2016-Datacenter",
              version: "latest",
            },
            osDisk: {
              caching: "ReadWrite",
              createOption: "FromImage",
              diskSizeGB: 512,
              managedDisk: { storageAccountType: "Standard_LRS" },
            },
          },
        },
      },
      sku: { name: "Standard_A1_v2", capacity: 2, tier: "Standard" },
      zones: ["1", "3"],
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}",
      subscriptionId,
      resourceGroupName,
      vmScaleSetName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAScaleSetWithVirtualMachinesInDifferentZones().catch(console.error);
