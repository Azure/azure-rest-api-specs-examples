const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithADiffOsDiskAndFullCachingEnabled.json
 */
async function createAVmWithEphemeralOsDiskAndEnableFullCachingSetToTrue() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    plan: { publisher: "microsoft-ads", product: "windows-data-science-vm", name: "windows2016" },
    hardwareProfile: { vmSize: "Standard_DS1_v2" },
    storageProfile: {
      imageReference: {
        sku: "windows2016",
        publisher: "microsoft-ads",
        version: "latest",
        offer: "windows-data-science-vm",
      },
      osDisk: {
        caching: "ReadOnly",
        diffDiskSettings: { option: "Local", placement: "TempDisk", enableFullCaching: true },
        managedDisk: { storageAccountType: "Standard_LRS" },
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
