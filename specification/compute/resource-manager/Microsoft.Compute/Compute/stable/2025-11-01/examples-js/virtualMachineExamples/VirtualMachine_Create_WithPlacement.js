const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithPlacement.json
 */
async function createAVMWithAutomaticZonePlacement() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus2",
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
    placement: { zonePlacementPolicy: "Any", includeZones: ["1", "3"] },
  });
  console.log(result);
}
