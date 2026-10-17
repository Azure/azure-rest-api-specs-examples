const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Create_WithFlexProfile.json
 */
async function createAVMWithAFlexProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    flexProfile: { preferredVMSize: "Standard_D4s_v5" },
    storageProfile: {
      imageReference: {
        publisher: "Canonical",
        offer: "0001-com-ubuntu-server-focal",
        sku: "20_04-lts-gen2",
        version: "latest",
      },
      osDisk: { createOption: "FromImage", managedDisk: { storageAccountType: "Premium_LRS" } },
    },
    osProfile: {
      computerName: "myVM",
      adminUsername: "azureuser",
      linuxConfiguration: {
        disablePasswordAuthentication: true,
        ssh: {
          publicKeys: [
            { path: "/home/azureuser/.ssh/authorized_keys", keyData: "ssh-rsa AAAA...{example}" },
          ],
        },
      },
    },
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/myNic",
          primary: true,
        },
      ],
    },
    tags: { department: "engineering", project: "alpha" },
  });
  console.log(result);
}
