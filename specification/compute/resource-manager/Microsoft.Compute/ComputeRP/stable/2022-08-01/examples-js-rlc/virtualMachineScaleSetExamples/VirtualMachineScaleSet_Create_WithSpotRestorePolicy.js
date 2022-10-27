const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSpotRestorePolicy.json
 */
async function createAScaleSetWithSpotRestorePolicy() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const options = {
    body: {
      location: "westus",
      properties: {
        overprovision: true,
        spotRestorePolicy: { enabled: true, restoreTimeout: "PT1H" },
        upgradePolicy: { mode: "Manual" },
        virtualMachineProfile: {
          billingProfile: { maxPrice: -1 },
          evictionPolicy: "Deallocate",
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
          priority: "Spot",
          storageProfile: {
            imageReference: {
              offer: "WindowsServer",
              publisher: "MicrosoftWindowsServer",
              sku: "2016-Datacenter",
              version: "latest",
            },
            osDisk: {
              caching: "ReadWrite",
              createOption: "FromImage",
              managedDisk: { storageAccountType: "Standard_LRS" },
            },
          },
        },
      },
      sku: { name: "Standard_A8m_v2", capacity: 2, tier: "Standard" },
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

createAScaleSetWithSpotRestorePolicy().catch(console.error);
