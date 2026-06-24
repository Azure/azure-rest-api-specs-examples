const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithSecurityTypeConfidentialVMWithCustomerManagedKeys.json
 */
async function createAVMWithSecurityTypeConfidentialVMWithCustomerManagedKeys() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_DC2as_v5" },
    securityProfile: {
      uefiSettings: { secureBootEnabled: true, vTpmEnabled: true },
      securityType: "ConfidentialVM",
    },
    storageProfile: {
      imageReference: {
        sku: "windows-cvm",
        publisher: "MicrosoftWindowsServer",
        version: "17763.2183.2109130127",
        offer: "2019-datacenter-cvm",
      },
      osDisk: {
        caching: "ReadOnly",
        managedDisk: {
          storageAccountType: "StandardSSD_LRS",
          securityProfile: {
            securityEncryptionType: "DiskWithVMGuestState",
            diskEncryptionSet: {
              id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
            },
          },
        },
        createOption: "FromImage",
        name: "myVMosdisk",
      },
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
