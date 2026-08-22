const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Create_WithAdditionalDiskProperties.json
 */
async function createAVmWithAdditionalDiskPropertiesNetworkAccessPolicyAndTier() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D4s_v3" },
    storageProfile: {
      imageReference: {
        sku: "2022-datacenter-azure-edition",
        publisher: "MicrosoftWindowsServer",
        version: "latest",
        offer: "WindowsServer",
      },
      diskApiVersion: "2026-03-02",
      osDisk: {
        caching: "ReadWrite",
        managedDisk: {
          storageAccountType: "Premium_LRS",
          additionalDiskProperties: {
            managedDiskProperties: {
              networkAccessPolicy: "AllowPrivate",
              diskAccessId:
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess",
            },
          },
        },
        name: "myVMosdisk",
        createOption: "FromImage",
      },
      dataDisks: [
        {
          lun: 0,
          name: "myDataDisk",
          createOption: "Empty",
          diskSizeGB: 1024,
          managedDisk: {
            storageAccountType: "Premium_LRS",
            additionalDiskProperties: {
              managedDiskProperties: {
                networkAccessPolicy: "AllowPrivate",
                diskAccessId:
                  "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess",
                tier: "P30",
              },
            },
          },
        },
      ],
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "myVM",
      adminPassword: "{your-password}",
    },
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
  });
  console.log(result);
}
