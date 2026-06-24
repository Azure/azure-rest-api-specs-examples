const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithInterconnectBlock.json
 */
async function createOrUpdateAVMWithInterconnectBlock() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_ND128isr_GB300_v6" },
    storageProfile: {
      imageReference: {
        publisher: "microsoft-dsvm",
        offer: "ubuntu-hpc",
        sku: "2404-gb",
        version: "latest",
      },
      osDisk: {
        caching: "ReadWrite",
        managedDisk: { storageAccountType: "Premium_LRS" },
        createOption: "FromImage",
        name: "myVMosdisk",
      },
    },
    interconnectBlockProfile: {
      interconnectBlock: {
        id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock",
      },
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "myVM",
      adminPassword: "{your-password}",
      linuxConfiguration: { disablePasswordAuthentication: false },
    },
    networkProfile: {
      interconnectGroupProfile: {
        interconnectGroup: {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup",
        },
        subgroups: [
          {
            id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0",
          },
        ],
      },
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
    zones: ["1"],
  });
  console.log(result);
}
