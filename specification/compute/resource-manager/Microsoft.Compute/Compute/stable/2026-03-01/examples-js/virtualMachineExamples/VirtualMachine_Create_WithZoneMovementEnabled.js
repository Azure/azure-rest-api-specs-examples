const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithZoneMovementEnabled.json
 */
async function createAVmWithResiliencyProfileAndZoneMovementEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D2s_v3" },
    storageProfile: {
      imageReference: {
        sku: "2016-Datacenter",
        publisher: "MicrosoftWindowsServer",
        version: "latest",
        offer: "WindowsServer",
      },
      osDisk: {
        caching: "ReadWrite",
        managedDisk: { storageAccountType: "Premium_ZRS" },
        name: "myVMosdisk",
        createOption: "FromImage",
      },
      dataDisks: [
        {
          diskSizeGB: 128,
          lun: 0,
          createOption: "Empty",
          managedDisk: { storageAccountType: "Premium_ZRS" },
          name: "myVMdatadisk1",
          caching: "ReadWrite",
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
    resiliencyProfile: { zoneMovement: { isEnabled: true } },
  });
  console.log(result);
}
