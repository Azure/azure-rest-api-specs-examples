const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithProtectedSettingsFromKeyVault.json
 */
async function createAVmssWithAnExtensionWithProtectedSettingsFromKeyVault() {
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
        upgradePolicy: { mode: "Manual" },
        virtualMachineProfile: {
          diagnosticsProfile: {
            bootDiagnostics: {
              enabled: true,
              storageUri: "http://{existing-storage-account-name}.blob.core.windows.net",
            },
          },
          extensionProfile: {
            extensions: [
              {
                name: "{extension-name}",
                properties: {
                  type: "{extension-Type}",
                  autoUpgradeMinorVersion: false,
                  protectedSettingsFromKeyVault: {
                    secretUrl:
                      "https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e",
                    sourceVault: {
                      id: "/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName",
                    },
                  },
                  publisher: "{extension-Publisher}",
                  settings: {},
                  typeHandlerVersion: "{handler-version}",
                },
              },
            ],
          },
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
      sku: { name: "Standard_D1_v2", capacity: 3, tier: "Standard" },
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

createAVmssWithAnExtensionWithProtectedSettingsFromKeyVault().catch(console.error);
